/*
 * Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package fields

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// Selector represents a field selector.
type Selector interface {
	// Matches returns true if this selector matches the given set of fields.
	Matches(Fields) bool

	// Empty returns true if this selector does not restrict the selection space.
	Empty() bool

	// RequiresExactMatch allows a caller to introspect whether a given selector
	// requires a single specific field to be set, and if so returns the value it
	// requires.
	RequiresExactMatch(field string) (value string, found bool)

	// Transform returns a new copy of the selector after TransformFunc has been
	// applied to the entire selector, or an error if fn returns an error.
	// If for a given requirement both field and value are transformed to empty
	// string, the requirement is skipped.
	Transform(fn TransformFunc) (Selector, error)

	// Requirements converts this interface to Requirements to expose
	// more detailed selection information.
	Requirements() Requirements

	// String returns a human-readable string that represents this selector.
	String() string

	// DeepCopySelector Make a deep copy of the selector.
	DeepCopySelector() Selector
}

type nothingSelector struct{}

func (n nothingSelector) Matches(_ Fields) bool      { return false }
func (n nothingSelector) Empty() bool                { return false }
func (n nothingSelector) String() string             { return "" }
func (n nothingSelector) Requirements() Requirements { return nil }
func (n nothingSelector) DeepCopySelector() Selector { return n }
func (n nothingSelector) RequiresExactMatch(field string) (value string, found bool) {
	return "", false
}
func (n nothingSelector) Transform(fn TransformFunc) (Selector, error) { return n, nil }

// Nothing return a Selector that matches no fields.
func Nothing() Selector {
	return nothingSelector{}
}

func Everything() Selector {
	return andTerm{}
}

type hasTerm struct {
	field, value string
}

func (t *hasTerm) Matches(f Fields) bool {
	return f.Get(t.field) == t.value
}

func (t *hasTerm) Empty() bool {
	return false
}

func (t *hasTerm) RequiresExactMatch(field string) (value string, found bool) {
	if t.field == field {
		return t.value, true
	}

	return "", false
}

func (t *hasTerm) Transform(fn TransformFunc) (Selector, error) {
	field, value, err := fn(t.field, t.value)
	if err != nil {
		return nil, err
	}
	if len(field) == 0 && len(value) == 0 {
		return Everything(), nil
	}

	return &hasTerm{field, value}, nil
}

func (t *hasTerm) Requirements() Requirements {
	return []Requirement{{
		Field:    t.field,
		Operator: Equals,
		Value:    t.value,
	}}
}

func (t *hasTerm) String() string {
	return fmt.Sprintf("%v=%v", t.field, EscapeValue(t.value))
}

func (t *hasTerm) DeepCopySelector() Selector {
	if t == nil {
		return nil
	}
	out := new(hasTerm)
	*out = *t

	return out
}

type notHasTerm struct {
	field, value string
}

func (t *notHasTerm) Matches(f Fields) bool {
	return f.Get(t.field) != t.value
}

func (t *notHasTerm) Empty() bool {
	return false
}

func (t *notHasTerm) RequiresExactMatch(field string) (value string, found bool) {
	return "", false
}

func (t *notHasTerm) Transform(fn TransformFunc) (Selector, error) {
	field, value, err := fn(t.field, t.value)
	if err != nil {
		return nil, err
	}
	if len(field) == 0 && len(value) == 0 {
		return Everything(), nil
	}

	return &notHasTerm{field, value}, nil
}

func (t *notHasTerm) Requirements() Requirements {
	return []Requirement{{
		Field:    t.field,
		Operator: NotEquals,
		Value:    t.value,
	}}
}

func (t *notHasTerm) String() string {
	return fmt.Sprintf("%v!=%v", t.field, EscapeValue(t.value))
}

func (t *notHasTerm) DeepCopySelector() Selector {
	if t == nil {
		return nil
	}
	out := new(notHasTerm)
	*out = *t

	return out
}

type andTerm []Selector

func (t andTerm) Matches(f Fields) bool {
	for _, q := range t {
		if !q.Matches(f) {
			return false
		}
	}

	return true
}

func (t andTerm) Empty() bool {
	if t == nil {
		return true
	}
	if len([]Selector(t)) == 0 {
		return true
	}
	for i := range t {
		if !t[i].Empty() {
			return false
		}
	}

	return true
}

func (t andTerm) RequiresExactMatch(field string) (string, bool) {
	if t == nil || len([]Selector(t)) == 0 {
		return "", false
	}
	for i := range t {
		if value, found := t[i].RequiresExactMatch(field); found {
			return value, found
		}
	}

	return "", false
}

func (t andTerm) Transform(fn TransformFunc) (Selector, error) {
	next := make([]Selector, 0, len([]Selector(t)))
	for _, s := range []Selector(t) {
		n, err := s.Transform(fn)
		if err != nil {
			return nil, err
		}
		if !n.Empty() {
			next = append(next, n)
		}
	}

	return andTerm(next), nil
}

func (t andTerm) Requirements() Requirements {
	reqs := make([]Requirement, 0, len(t))
	for _, s := range []Selector(t) {
		rs := s.Requirements()
		reqs = append(reqs, rs...)
	}

	return reqs
}

func (t andTerm) String() string {
	terms := make([]string, 0, len(t))
	for _, q := range t {
		terms = append(terms, q.String())
	}

	return strings.Join(terms, ",")
}

func (t andTerm) DeepCopySelector() Selector {
	if t == nil {
		return nil
	}
	out := make([]Selector, 0, len(t))
	for i := range t {
		out[i] = t[i].DeepCopySelector()
	}

	return andTerm(out)
}

func SelectorFromSet(s Set) Selector {
	if s == nil {
		return Everything()
	}

	fields := make([]Selector, 0, len(s))
	for field, value := range s {
		fields = append(fields, &hasTerm{field: field, value: value})
	}
	if len(fields) == 1 {
		return fields[0]
	}

	return andTerm(fields)
}

var valueEscaper = strings.NewReplacer(
	`\`, `\\`,
	`,`, `\,`,
	`=`, `\=`,
)

func EscapeValue(s string) string {
	return valueEscaper.Replace(s)
}

type InvalidEscapeSequence struct {
	sequence string
}

func (i InvalidEscapeSequence) Error() string {
	return fmt.Sprintf("invalid field selector: invalid escape sequence: %s", i.sequence)
}

type UnescapedRune struct {
	r rune
}

func (u UnescapedRune) Error() string {
	return fmt.Sprintf("invalid field selector: unescapedRune character in value: %v", u.r)
}

func UnescapeValue(s string) (string, error) {
	if !strings.ContainsAny(s, `\,=`) {
		return s, nil
	}

	buf := bytes.NewBuffer(make([]byte, 0, len(s)))
	inSlash := false
	for _, c := range s {
		if inSlash {
			switch c {
			case '\\', ',', '=':
				// omit the \ for recognized escape sequences
				buf.WriteRune(c)
			default:
				// error on unrecognized escape sequences
				return "", InvalidEscapeSequence{sequence: string([]rune{'\\', c})}
			}
			inSlash = false

			continue
		}

		switch c {
		case '\\':
			inSlash = true
		case ',', '=':
			// unescaped , and = characters are not allowed in field selector values
			return "", UnescapedRune{r: c}
		default:
			buf.WriteRune(c)
		}
	}

	if inSlash {
		return "", InvalidEscapeSequence{sequence: "\\"}
	}

	return buf.String(), nil
}

func ParseSelector(selector string) (Selector, error) {
	return parseSelector(selector,
		func(lhs, rhs string) (newLhs, newRhs string, err error) {
			return lhs, rhs, nil
		})
}

// ParseAndTransformSelector parses the selector and runs them through the given TransformFunc.
func ParseAndTransformSelector(selector string, fn TransformFunc) (Selector, error) {
	return parseSelector(selector, fn)
}

// TransformFunc transforms selectors.
type TransformFunc func(field, value string) (newField, newValue string, err error)

// splitTerms returns the comma-separated terms contained in the given fieldSelector.
// Backslash-escaped commas are treated as data instead of delimiters,
// and are included in the returned terms, with the leading backslash preserved.
func splitTerms(fieldSelector string) []string {
	if len(fieldSelector) == 0 {
		return nil
	}

	terms := make([]string, 0, 1)
	startIndex := 0
	inSlash := false
	for i, c := range fieldSelector {
		switch {
		case inSlash:
			inSlash = false
		case c == '\\':
			inSlash = true
		case c == ',':
			terms = append(terms, fieldSelector[startIndex:i])
			startIndex = i + 1
		}
	}

	terms = append(terms, fieldSelector[startIndex:])

	return terms
}

const (
	notEqualOperator    = "!="
	doubleEqualOperator = "=="
	equalOperator       = "="
)

// termOperators holds the recognized operators supported in fieldSelectors.
// doubleEqualOperator and equal are equivalent, but doubleEqualOperator is checked first
// to avoid leaving a leading = character on the rhs value.
var termOperators = []string{notEqualOperator, doubleEqualOperator, equalOperator}

// splitTerm returns the lhs, operator, and rhs parsed from the given term, along with an
// indicator of whether the parse was successful.
// No escaping of special characters is supported in the lhs value,
// so the first occurrence of a recognized operator is used as the split point.
// The literal rhs is returned, and the caller is responsible for applying any desired uncapping.
func splitTerm(term string) (lhs, op, rhs string, ok bool) {
	for i := range term {
		remaining := term[i:]
		for _, op := range termOperators {
			if strings.HasPrefix(remaining, op) {
				return term[0:i], op, term[i+len(op):], true
			}
		}
	}

	return "", "", "", false
}

func parseSelector(selector string, fn TransformFunc) (Selector, error) {
	parts := splitTerms(selector)
	sort.StringSlice(parts).Sort()
	var items []Selector
	for _, part := range parts {
		if part == "" {
			continue
		}
		lhs, op, rhs, ok := splitTerm(part)
		if !ok {
			return nil, fmt.Errorf("invalid selector: '%s'; can't understand '%s'", selector, part)
		}
		unescapedRHS, err := UnescapeValue(rhs)
		if err != nil {
			return nil, err
		}
		switch op {
		case notEqualOperator:
			items = append(items, &notHasTerm{field: lhs, value: unescapedRHS})
		case doubleEqualOperator:
			items = append(items, &hasTerm{field: lhs, value: unescapedRHS})
		case equalOperator:
			items = append(items, &hasTerm{field: lhs, value: unescapedRHS})
		default:
			return nil, fmt.Errorf("invalid selector: '%s'; can't understand '%s'", selector, part)
		}
	}
	if len(items) == 1 {
		return items[0].Transform(fn)
	}

	return andTerm(items).Transform(fn)
}

// OneTermEqualSelector returns an object that matches objects where one field/field equals one value.
// Cannot return an error.
func OneTermEqualSelector(k, v string) Selector {
	return &hasTerm{field: k, value: v}
}

// OneTermNotEqualSelector returns an object that matches objects where one field/field does not equal one value.
// Cannot return an error.
func OneTermNotEqualSelector(k, v string) Selector {
	return &notHasTerm{field: k, value: v}
}

// AndSelectors create a selector that is the logical AND of all the given selectors.
func AndSelectors(selectors ...Selector) Selector {
	return andTerm(selectors)
}

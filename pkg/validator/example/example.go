/*
 * Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"log"

	"blog-api/pkg/validator"
	go_validator "github.com/go-playground/validator/v10"
)

type Data struct {
	Name string `json:"name" validate:"required,gt=0,lte=4" label:"姓名" label_en:"name"`
	Age  uint   `json:"age" validate:"required,gte=0,lte=50" label:"年龄" label_en:"age"`
	Sex  uint   `json:"sex" validate:"required,oneof=1 2 3" label:"性别" label_en:"sex"`
}

type Data2 struct {
	Age uint `json:"age" validate:"required,myValidation,gte=0,lte=50" label:"年龄" label_en:"age"`
}

func main() {
	d1 := &Data{
		Age: 66, // 不符合验证规则
		Sex: 1,
	}

	opts := &validator.Options{
		Language: "zh",
		Tag:      "label",
	}
	validator.Init(opts)

	if err := validator.Struct(d1); err != nil {
		// output: [姓名为必填字段 年龄必须小于或等于50]
		fmt.Println(err.(*validator.ValidationErrors).TranslateErrs())
	}

	d2 := &Data{
		Name: "jack",
		Age:  51,
	}
	v1 := validator.New("en", "label_en")
	if err := v1.Struct(d2); err != nil {
		// output: map[age:age must be 50 or less sex:sex is a required field]
		fmt.Println(err.(*validator.ValidationErrors).TranslateErrsMap())
	}

	d3 := &Data2{Age: 66}
	v2 := validator.New("zh", "label")
	// register custom validation.
	if rErr := v2.RegisterValidation("myValidation", "{0}真的不能为66", myValidation); rErr != nil {
		log.Fatalln(rErr.Error())
	}
	if err := v2.Struct(d3); err != nil {
		fmt.Println(err.(*validator.ValidationErrors).TranslateErrs())    // output: [年龄真的不能为66]
		fmt.Println(err.Error())                                          // output: 年龄真的不能为66
		fmt.Println(err.(*validator.ValidationErrors).TranslateErrsMap()) // output: map[年龄:年龄真的不能为66]
	}

	d4 := &Data{Name: "tom", Age: 10, Sex: 2}
	if err := validator.Struct(d4); err != nil {
		fmt.Println(err.(*validator.ValidationErrors).TranslateErrs())
	} else {
		fmt.Println("success") // output: success
	}
}

func myValidation(fl go_validator.FieldLevel) bool {
	if fl.Field().Uint() == 66 {
		return false
	}

	return true
}

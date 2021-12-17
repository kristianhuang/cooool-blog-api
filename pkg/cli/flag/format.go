package flag

import (
	"github.com/spf13/pflag"
	"log"
	"strings"
)

// WordReplaceNormalizeFunc 转化非标准格式 flag 为标准格式
func WordReplaceNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

// WarnWordSepNormalizeFunc changes and warns for flags that contain "_" separators.
func WarnWordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		name := strings.ReplaceAll(name, "_", "-")
		log.Printf("%s is DEPRECATED and will be removed in a future version. Use %s instead.", name, name)
		return pflag.NormalizedName(name)
	}
	return pflag.NormalizedName(name)
}

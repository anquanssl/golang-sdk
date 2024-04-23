package utils

import (
    "fmt"
    "reflect"
    "sort"
    "strings"
)

func HttpBuildQuery(data map[string]interface{}, prefix string) string {
    if data == nil || reflect.TypeOf(data).Kind() != reflect.Map {
        return ""
    }

    encodeURIComponent := func(str string) string {
        m := map[string]string{
            " ":  "+",
            "\n": "%0A",
            "!":  "%21",
            "\"": "%22",
            "#":  "%23",
            "$":  "%24",
            "%":  "%25",
            "&":  "%26",
            "'":  "%27",
            "(":  "%28",
            ")":  "%29",
            "*":  "%2A",
            "+":  "%2B",
            ",":  "%2C",
            "/":  "%2F",
            ":":  "%3A",
            ";":  "%3B",
            "<":  "%3C",
            "=":  "%3D",
            ">":  "%3E",
            "?":  "%3F",
            "@":  "%40",
            "[":  "%5B",
            "\\": "%5C",
            "]":  "%5D",
            "^":  "%5E",
            "`":  "%60",
            "{":  "%7B",
            "|":  "%7C",
            "}":  "%7D",
            "~":  "%7E",
        }

        result := ""
        for _, c := range str {
            if val, ok := m[string(c)]; ok {
                result += val
            } else {
                result += string(c)
            }
        }
        return result
    }

    var query []string
    keys := make([]string, 0, len(data))
    for k := range data {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    for _, key := range keys {
        param := key
        value := data[key]
        if prefix != "" {
            key = prefix + "[" + param + "]"
        }

        switch v := value.(type) {
        case map[string]interface{}:
            query = append(query, HttpBuildQuery(v, key))
        default:
            query = append(query, encodeURIComponent(key)+"="+encodeURIComponent(fmt.Sprintf("%v", value)))
        }
    }

    return strings.Join(query, "&")
}

package utils

import "strings"

func SplitCSVStringIntoFields(line string) ([]string, error) {

	var fields []string

	raw := strings.Split(line, ",")

	for i := 0; i < len(raw); i++ {
		field := strings.Trim(raw[i], " ")

		if strings.HasPrefix(field, "\"") {
			// need to find the last field, which will end with a quote
			field = field[1:]
			if strings.HasSuffix(field, "\"") {
				field = field[:len(field)-1]
			} else {
				for i++; i < len(raw); i++ {
					subfield := raw[i]
					field += "," + subfield
					if strings.HasSuffix(subfield, "\"") {
						field = field[:len(field)-1]
						break
					}
				}
			}
		}
		fields = append(fields, field)
	}

	return fields, nil
}

func SplitMultilineCSV(raw string, skipFirst bool) ([][]string, error) {

	raw = strings.Replace(raw, "\r", "", -1)
	lines := strings.Split(raw, "\n")
	if skipFirst {
		lines = lines[1:]
	}

	parsed := make([][]string, len(lines))

	for i := 0; i < len(lines); i++ {
		fields, err := SplitCSVStringIntoFields(lines[i])
		if err != nil {
			return parsed, err
		}
		parsed[i] = fields
	}

	return parsed, nil

}

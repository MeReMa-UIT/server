package utils

import "strings"

func GetAttachmentPrefix(filename string) string {
	var prefixes = []string{"xray_", "ct_", "ultrasound_", "test_", "other_"}

	for _, prefix := range prefixes {
		if strings.HasPrefix(filename, prefix) {
			return prefix[0 : len(prefix)-1]
		}
	}
	return ""
}

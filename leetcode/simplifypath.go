import (
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{}

	chunks := strings.Split(path, "/")

	for _, chunk := range chunks {

		switch chunk {
		case "", ".":
			continue
		case "..":
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, chunk)

		}

	}

	return "/" + strings.Join(stack, "/")

}
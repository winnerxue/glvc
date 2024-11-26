package strings

import "bytes"

// 定义一个 Builder 类型，模仿 strings.Builder 的接口
type Builder struct {
	buf bytes.Buffer
}

func (sb *Builder) WriteString(s string) (int, error) {
	return sb.buf.WriteString(s)
}

func (sb *Builder) WriteByte(c byte) error {
	return sb.buf.WriteByte(c)
}

func (sb *Builder) WriteRune(r rune) (int, error) {
	return sb.buf.WriteRune(r)
}

func (sb *Builder) Write(p []byte) (int, error) {
	return sb.buf.Write(p)
}

func (sb *Builder) String() string {
	return sb.buf.String()
}

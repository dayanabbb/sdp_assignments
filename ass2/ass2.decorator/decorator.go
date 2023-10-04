package main

import "fmt"

type TextFormatter interface {
	Format(text string) string
}

type UsualTextFormatter struct{}

func (d *UsualTextFormatter) Format(text string) string {
	return text
}

type Uppercase struct {
	formatter TextFormatter
}

func (u *Uppercase) Format(text string) string {
	return fmt.Sprintf("[Uppercase] %s", u.formatter.Format(text))
}

type Bold struct {
	formatter TextFormatter
}

func (b *Bold) Format(text string) string {
	return fmt.Sprintf("[Bold] %s", b.formatter.Format(text))
}

func main() {

	formatter := &UsualTextFormatter{}

	decoratedFormatter := &Bold{
		formatter: &Uppercase{
			formatter: formatter,
		},
	}

	formattedText := decoratedFormatter.Format("dayana belyalova")
	fmt.Println(formattedText)
}

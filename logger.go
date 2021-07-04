package logger

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"path"
	"time"
)

var logLocation string

func Fatal(format string, a ...interface{}) {
	write(format, a...)
	_, _ = fmt.Fprintf(color.Output,"%s %s\n", header("fatal"), color.RedString(format, a...))
	os.Exit(1)
}

func Error(format string, a ...interface{}) {
	write(format, a...)
	_, _ = fmt.Fprintf(color.Output,"%s %s\n", header("error"), color.RedString(format, a...))
}

func Warning(format string, a ...interface{}) {
	write(format, a...)
	warning(format, a...)
}

func Info(format string, a ...interface{}) {
	write(format, a...)
	_, _ = fmt.Fprintf(color.Output, "%v %v\n", header("info"), color.CyanString(format, a...))
}

func Success(format string, a ...interface{}) {
	write(format, a...)
	_, _ = fmt.Fprintf(color.Output,"%s %s\n", header("success"), color.GreenString(format, a...))
}

func Log(format string, a ...interface{}) {
	write(format, a...)
	_, _ = fmt.Fprintf(color.Output,"%s %s\n", header("log"), fmt.Sprintf(format, a...))
}

func SetLogLocation(path string) {
	d, err := os.Stat(path)
	if err == nil && !d.IsDir() {
		log.Fatalf("%s already exist but isn't a directory", path)
	}

	if os.IsNotExist(err) {
		err = os.Mkdir(path, 0700)
		if err != nil && err != os.ErrExist {
			log.Fatal(err)
		}
	}

	logLocation = path
}

func header(t string) string {
	return color.MagentaString("[%s][%s]", time.Now().Format("15:04:05"), t)
}

func write(format string, a ...interface{}) {
	if logLocation != "" {
		t := time.Now()
		p := path.Join(
			logLocation,
			fmt.Sprintf("%s.log", t.Format("20060102")),
		)

		f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			warning("problem has occured, cannot write into %s", p)
			return
		}
		defer f.Close()

		_, err = f.WriteString(fmt.Sprintf(
			"[%s] %s",
			t.Format("2006-01-02 15:04:05"),
			fmt.Sprintf(format, a...),
			))
		if err != nil {
			warning("problem has occured, cannot write into %s", p)
			return
		}
	}
}

func warning(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(color.Output,"%s %s\n", header("warning"), color.YellowString(format, a...))
}
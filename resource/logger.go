package resource

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const (
	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite

	colorBold     = 1
	colorDarkGray = 90
)

var Logger *Loggers

type Loggers struct {
	zLog zerolog.Logger
}

func InitLogger() {
	if Logger == nil {
		Logger = new(Loggers)
		if viper.GetString("env") == "online" {
			Logger.zLog.Level(zerolog.InfoLevel)
			Logger.zLog = zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
				w.NoColor = true
				w.TimeFormat = "[2006-01-02 15:04:05]"
				w.FormatLevel = func(i interface{}) string {
					return strings.ToUpper(fmt.Sprintf("[%s]", i))
				}
				w.FormatCaller = func(i interface{}) string {
					return fmt.Sprintf("%s", i)
				}
				w.Out = os.Stdout
			})).With().Timestamp().CallerWithSkipFrameCount(3).Logger()
		} else {
			Logger.zLog.Level(zerolog.DebugLevel)
			Logger.zLog = zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
				w.Out = os.Stdout
			})).With().Timestamp().CallerWithSkipFrameCount(3).Logger()
		}
	}
}

func (l *Loggers) Output() (w io.Writer) {
	return l.zLog.Output(w)
}
func (l *Loggers) SetOutput(w io.Writer) {
	l.zLog.Output(w)
}
func (l *Loggers) Prefix() string {
	return ""
}
func (l *Loggers) SetPrefix(p string) {
	return
}
func (l *Loggers) Level() log.Lvl {
	return log.Lvl(l.zLog.GetLevel())
}
func (l *Loggers) SetLevel(v log.Lvl) {
	l.zLog.Level(zerolog.Level(v))
}
func (l *Loggers) SetHeader(h string) {
	return
}
func (l *Loggers) Print(i ...interface{}) {
	l.zLog.Info().Msg(fmt.Sprint(i))
}
func (l *Loggers) Printf(format string, args ...interface{}) {
	l.zLog.Info().Msgf(format, args...)
}
func (l *Loggers) Printj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.zLog.Info().Msg(string(b))
}
func (l *Loggers) Debug(i ...interface{}) {
	l.zLog.Debug().Msg(fmt.Sprint(i...))
}
func (l *Loggers) Debugf(format string, args ...interface{}) {
	l.zLog.Debug().Msgf(format, args...)
}
func (l *Loggers) Debugj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.zLog.Debug().Msg(string(b))
}
func (l *Loggers) Info(i ...interface{}) {
	l.zLog.Info().Msg(fmt.Sprint(i...))
}
func (l *Loggers) Infof(format string, args ...interface{}) {
	l.zLog.Info().Msgf(format, args...)
}
func (l *Loggers) Infoj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.zLog.Info().Msg(string(b))
}
func (l *Loggers) Warn(i ...interface{}) {
	l.zLog.Warn().Msg(fmt.Sprint(i...))
}
func (l *Loggers) Warnf(format string, args ...interface{}) {
	l.zLog.Warn().Msgf(format, args...)
}
func (l *Loggers) Warnj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.zLog.Warn().Msg(string(b))
}
func (l *Loggers) Error(i ...interface{}) {
	l.zLog.Error().Msg(fmt.Sprint(i...))
}
func (l *Loggers) Errorf(format string, args ...interface{}) {
	l.zLog.Error().Msgf(format, args...)
}
func (l *Loggers) Errorj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.zLog.Error().Msg(string(b))
}
func (l *Loggers) Fatal(i ...interface{}) {
	l.zLog.Fatal().Msg(fmt.Sprint(i...))
}
func (l *Loggers) Fatalj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.zLog.Fatal().Msg(string(b))
}
func (l *Loggers) Fatalf(format string, args ...interface{}) {
	l.zLog.Fatal().Msgf(format, args...)
}
func (l *Loggers) Panic(i ...interface{}) {
	l.zLog.Panic().Msg(fmt.Sprint(i...))
}
func (l *Loggers) Panicj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.zLog.Panic().Msg(string(b))
}
func (l *Loggers) Panicf(format string, args ...interface{}) {
	l.zLog.Panic().Msgf(format, args...)
}

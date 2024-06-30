package config

import (
	"io"
	"os"
	"path"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config for logging
type Config struct {
	// ConsoleLoggingEnabled Enable console logging
	ConsoleLoggingEnabled bool
	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

type Logger struct {
	*zerolog.Logger
}

func newRollingFile(config Config) io.Writer {
	// if err := os.MkdirAll(config.Directory, 0744); err != nil {
	// 	log.Error().Err(err).Str("path", config.Directory).Msg("can't create log directory")
	// 	return nil
	// }

	return &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxBackups: config.MaxBackups, // files
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
	}
}

// Configure sets up the logging framework
//
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output log file will be located at /var/log/service-xyz/service-xyz.log and
// will be rolled according to configuration set.
func Configure(config Config) *Logger {
	var writers []io.Writer

	if config.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}
	mw := io.MultiWriter(writers...)

	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", config.FileLoggingEnabled).
		Bool("jsonLogOutput", config.EncodeLogsAsJson).
		Str("logDirectory", config.Directory).
		Str("fileName", config.Filename).
		Int("maxSizeMB", config.MaxSize).
		Int("maxBackups", config.MaxBackups).
		Int("maxAgeInDays", config.MaxAge).
		Msg("logging configured")

	return &Logger{
		Logger: &logger,
	}
}

func logWithConfig() *Logger {
	//logfileMaxSize, err := strconv.Atoi(config.LogfileMaxSize)
	//if err != nil {
	//	fmt.Printf("err %v\n", err)
	//}

	//configLog merge to main config not yet implemented because of importing loop
	configLog := Config{
		ConsoleLoggingEnabled: true,
		EncodeLogsAsJson:      true,
		FileLoggingEnabled:    true,
		Directory:             "log",
		Filename:              "tkbai.log",
		MaxSize:               2,
	}
	return Configure(configLog)
}

var Log = *logWithConfig()

func LogErr(err error, msg string) {
	Log.Error().Err(err).Msg(msg)
}

func LogTrc(funcName, msg string) {
	Log.Trace().Str("FUNC", funcName).Msg(msg)
}

func LogDbg(funcName, msg string) {
	Log.Debug().Str("FUNC", funcName).Msg(msg)
}

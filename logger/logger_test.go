package logger_test

// Acceptance testing for groot. Groot
// library provides abstraction, they are tested here.
// If you are adding more test cases, add those in Makefile under "test-acceptance" target.

// WARNING: Do not run these tests together. Run these tests individually
// using `make test-acceptance` command.
// If you run it as a whole, it invokes logger sinks from internal(zap) init functions and
// it fails by repeatedly trying to add more syncs, as packages are already imported by main routine.

// XXX: Think about abstracting the sinks to run tests effectively.
// https://groups.google.com/g/golang-nuts/c/lrVU58AMI4c/m/Wfn51pUTtMkJ

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/arcana-network/groot/logger"
	"github.com/stretchr/testify/require"
)

func TestNewZapLogger(t *testing.T) {
	t.Parallel()

	name := "withServiceName"
	service := "test"

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		require.NotPanics(t, func() { logger.NewZapLogger(service) })
	})
}

func TestNewGlobalZapLogger(t *testing.T) {
	t.Parallel()

	name := "withServiceName"
	logger.NewZapGlobal(name)

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		// Won't panic because of registered file sinks how many ever time you import
		require.NotPanics(t, func() { logger.NewZapGlobal(name) })
	})
}

func TestNewZapLoggerEmptyService(t *testing.T) {
	t.Parallel()

	name := "withServiceName"
	service := ""
	expectedErr := "service cannot be empty"

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		require.PanicsWithError(t, expectedErr, func() { logger.NewZapLogger(service) })
	})
}

func TestNewGlobalZapLoggerEmptyService(t *testing.T) {
	t.Parallel()

	name := "withServiceName"
	service := ""
	expectedErr := "service cannot be empty"

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		require.PanicsWithError(t, expectedErr, func() { logger.NewZapGlobal(service) })
	})
}

func TestSinkRepeat(t *testing.T) {
	t.Parallel()

	name := "sinkRepeat"
	expectedErr := `register lumberjack sync: sink factory already registered for scheme "lumberjack"`

	// Create a new logger, it registers sinks.
	logger.NewZapLogger(name)

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		// By creating one more logger from same instance, it panics by registering the same sink.
		require.PanicsWithError(t, expectedErr, func() { logger.NewZapLogger(name) })
	})
}

func TestGlobalSinkRepeat(t *testing.T) {
	t.Parallel()

	name := "sinkRepeatGlobal"

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		require.NotPanics(t, func() { logger.NewZapGlobal(name) })
		// initializing global loggger multiple times shouldn't panic
		require.NotPanics(t, func() { logger.NewZapGlobal(name) })
		require.NotPanics(t, func() { logger.NewZapGlobal(name) })
		require.NotPanics(t, func() { logger.NewZapGlobal(name) })
	})
}

func TestFileCreation(t *testing.T) {
	t.Parallel()

	home, err := os.UserHomeDir()
	if err != nil {
		t.FailNow()
	}

	name := "fileCreation"
	fname := fmt.Sprintf("/%s.log", name)

	wantFile := path.Join(home, "/arcana/logs", name, fname)

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		log := logger.NewZapLogger(name)
		log.Info("Test info 1", logger.Field{
			"test key 1": "test value 1",
			"test key 2": "test value 2",
		})
		log.Info("Test info 1", logger.Field{
			"test key 3": "test value 3",
			"test key 4": "test value 4",
		})
		_, err := os.Open(wantFile)
		require.NoError(t, err)
	})

	t.Cleanup(func() {
		cleanupFile(wantFile)
	})
}

func TestGlobalFileCreation(t *testing.T) {
	t.Parallel()

	home, err := os.UserHomeDir()
	if err != nil {
		t.FailNow()
	}

	name := "fileCreationGlobal"
	fname := fmt.Sprintf("/%s.log", name)

	wantFile := path.Join(home, "/arcana/logs", name, fname)

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		log := logger.NewZapGlobal(name)
		log.Info("Test info 1", logger.Field{
			"test key 1": "test value 1",
			"test key 2": "test value 2",
		})
		log = logger.NewZapGlobal(name) // initializing global loggger multiple times shouldn't change anything
		log.Info("Test info 1", logger.Field{
			"test key 3": "test value 3",
			"test key 4": "test value 4",
		})
		_, err := os.Open(wantFile)
		require.NoError(t, err)
	})

	t.Cleanup(func() {
		cleanupFile(wantFile)
	})
}

func TestFileContent(t *testing.T) {
	t.Parallel()

	home, err := os.UserHomeDir()
	if err != nil {
		t.FailNow()
	}

	name := "fileContent"
	fname := fmt.Sprintf("/%s.log", name)
	wantFile := path.Join(home, "/arcana/logs", name, fname)
	gotLines := 0
	wantLines := 4

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		log := logger.NewZapLogger(name)
		log.Info("Test info 1", logger.Field{
			"info key 1": "info value 1",
			"info key 2": "info value 2",
		})
		log.Debug("Test debug 1", logger.Field{
			"debug key 3": "debug value 3",
			"debug key 4": "debug value 4",
		})
		log.Warn("Test warn 1", logger.Field{
			"warn key 3": "warn value 3",
			"warn key 4": "warn value 4",
		})
		log.Error("Test error 1", logger.Field{
			"error key 3": "error value 3",
			"error key 4": "error value 4",
		})
		f, err := os.Open(wantFile)
		require.NoError(t, err)
		fb, err := ioutil.ReadAll(f)
		require.NoError(t, err)
		gotLines += bytes.Count(fb, []byte{'\n'})
		require.Equal(t, wantLines, gotLines) // proves we have given logs written in the file
	})

	t.Cleanup(func() {
		cleanupFile(wantFile)
	})
}

func TestGlobalFileContent(t *testing.T) {
	t.Parallel()

	home, err := os.UserHomeDir()
	if err != nil {
		t.FailNow()
	}

	name := "fileContentGlobal"
	fname := fmt.Sprintf("/%s.log", name)
	wantFile := path.Join(home, "/arcana/logs", name, fname)
	gotLines := 0
	wantLines := 4

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		log := logger.NewZapGlobal(name)
		log.Info("Test info 1", logger.Field{
			"info key 1": "info value 1",
			"info key 2": "info value 2",
		})
		log.Debug("Test debug 1", logger.Field{
			"debug key 3": "debug value 3",
			"debug key 4": "debug value 4",
		})

		log = logger.NewZapGlobal(name) // initializing global loggger multiple times shouldn't change anything

		log.Warn("Test warn 1", logger.Field{
			"warn key 3": "warn value 3",
			"warn key 4": "warn value 4",
		})
		log.Error("Test error 1", logger.Field{
			"error key 3": "error value 3",
			"error key 4": "error value 4",
		})
		f, err := os.Open(wantFile)
		require.NoError(t, err)
		fb, err := ioutil.ReadAll(f)
		require.NoError(t, err)
		gotLines += bytes.Count(fb, []byte{'\n'})
		require.Equal(t, wantLines, gotLines) // proves we have given logs written in the file
	})

	t.Cleanup(func() {
		cleanupFile(wantFile)
	})
}

func TestFatalLogs(t *testing.T) {
	t.Parallel()

	home, err := os.UserHomeDir()
	if err != nil {
		t.FailNow()
	}

	name := "fatalLog"
	fname := fmt.Sprintf("/%s.log", name)
	wantFile := path.Join(home, "/arcana/logs", name, fname)
	gotLines := 0
	wantLines := 5

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		// Fatal log does os.exit, let it run in different process and return
		// back to previous callsite
		if os.Getenv("FATAL_LOG") == "1" {
			log := logger.NewZapLogger(name)
			log.Info("Test info 1", logger.Field{
				"info key 1": "info value 1",
				"info key 2": "info value 2",
			})
			log.Debug("Test debug 1", logger.Field{
				"debug key 3": "debug value 3",
				"debug key 4": "debug value 4",
			})
			log.Warn("Test warn 1", logger.Field{
				"warn key 3": "warn value 3",
				"warn key 4": "warn value 4",
			})
			log.Error("Test error 1", logger.Field{
				"error key 3": "error value 3",
				"error key 4": "error value 4",
			})
			log.Fatal("Test fatal 1", logger.Field{
				"fatal key 3": "error value 3",
				"fatal key 4": "error value 4",
			})

			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestFatalLogs") //nolint // We have to launch sub process to test os.exit(fatal).
		cmd.Env = append(os.Environ(), "FATAL_LOG=1")
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() { // nolint:errorlint // type assert to access Success method
			f, err := os.Open(wantFile)
			require.NoError(t, err)
			fb, err := ioutil.ReadAll(f)
			require.NoError(t, err)
			gotLines += bytes.Count(fb, []byte{'\n'})
			require.Equal(t, wantLines, gotLines) // proves we have logs written in the file even after crashing
			t.Cleanup(func() {
				cleanupFile(wantFile)
			})

			return
		}
	})
}

func TestGlobalFatalLogs(t *testing.T) {
	t.Parallel()

	home, err := os.UserHomeDir()
	if err != nil {
		t.FailNow()
	}

	name := "fatalLogGlobal"
	fname := fmt.Sprintf("/%s.log", name)
	wantFile := path.Join(home, "/arcana/logs", name, fname)
	gotLines := 0
	wantLines := 5

	t.Run(name, func(t *testing.T) {
		t.Parallel()
		// Fatal log does os.exit, let it run in different process and return
		// back to previous callsite
		if os.Getenv("FATAL_LOG") == "1" {
			log := logger.NewZapGlobal(name)
			log.Info("Test info 1", logger.Field{
				"info key 1": "info value 1",
				"info key 2": "info value 2",
			})
			log.Debug("Test debug 1", logger.Field{
				"debug key 3": "debug value 3",
				"debug key 4": "debug value 4",
			})
			log = logger.NewZapGlobal(name) // initializing global loggger multiple times shouldn't change anything
			log.Warn("Test warn 1", logger.Field{
				"warn key 3": "warn value 3",
				"warn key 4": "warn value 4",
			})
			log.Error("Test error 1", logger.Field{
				"error key 3": "error value 3",
				"error key 4": "error value 4",
			})
			log.Fatal("Test fatal 1", logger.Field{
				"fatal key 3": "error value 3",
				"fatal key 4": "error value 4",
			})

			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestFatalLogsGlobal") //nolint // We have to launch sub process to test os.exit(fatal).
		cmd.Env = append(os.Environ(), "FATAL_LOG=1")
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() { // nolint:errorlint // type assert to access Success method
			f, err := os.Open(wantFile)
			require.NoError(t, err)
			fb, err := ioutil.ReadAll(f)
			require.NoError(t, err)
			gotLines += bytes.Count(fb, []byte{'\n'})
			require.Equal(t, wantLines, gotLines) // proves we have logs written in the file even after crashing
			t.Cleanup(func() {
				cleanupFile(wantFile)
			})

			return
		}
	})
}

// Helper function to cleanup files created by tests.
func cleanupFile(f string) {
	os.Remove(f)
}

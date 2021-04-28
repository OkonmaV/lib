package lib

import (
	"github.com/thin-peak/httpservice"

	"github.com/thin-peak/logger"
)

func LogsInit(configurator *httpservice.Configurator) ([]logger.LogWriter, error) {
	logsAddr, err := configurator.GetServiceAddress(httpservice.ServiceName(ServiceNameLogs))
	if err != nil {
		return nil, err
	}

	logsHttpWriter, err := logger.NewHTTPLogWriter(logsAddr[0], logger.DebugLevel)
	if err != nil {
		return nil, err
	}

	return []logger.LogWriter{logger.NewConsoleLogWriter(logger.DebugLevel), logsHttpWriter}, nil
}

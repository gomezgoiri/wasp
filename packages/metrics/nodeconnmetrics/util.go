package nodeconnmetrics

import (
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	chainLabelNameConst   = "chain"
	msgTypeLabelNameConst = "message_type"
)

func getMetricsLabel(chainID *iscp.ChainID, msgType string) prometheus.Labels {
	var chainIDStr string
	if chainID == nil {
		chainIDStr = ""
	} else {
		chainIDStr = chainID.String()
	}
	return prometheus.Labels{
		chainLabelNameConst:   chainIDStr,
		msgTypeLabelNameConst: msgType,
	}
}
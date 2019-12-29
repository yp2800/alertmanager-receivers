package alerts

import "time"

type AlertRecord struct {
	Receiver string `json:"receiver"`
	Status string `json:"status"`
	Alerts []struct {
		Status string `json:"status"`
		Labels struct {
			Alertname string `json:"alertname,omitempty"`
			Severity string `json:"severity,omitempty"`
			App string `json:"app,omitempty"`
			Chart string `json:"chart,omitempty"`
			Component string `json:"component,omitempty"`
			Heritage string `json:"heritage,omitempty"`
			Instance string `json:"instance,omitempty"`
			Job string `json:"job,omitempty"`
			KubernetesName string `json:"kubernetes_name,omitempty"`
			KubernetesNamespace string `json:"kubernetes_namespace,omitempty"`
			KubernetesNode string `json:"kubernetes_node,omitempty"`
			Release string `json:"release,omitempty"`
		} `json:"labels"`
		Annotations struct {
			Description string `json:"description"`
			Summary string `json:"summary"`
		} `json:"annotations"`
		StartsAt time.Time `json:"startsAt"`
		EndsAt time.Time `json:"endsAt"`
		GeneratorURL string `json:"generatorURL"`
		Fingerprint string `json:"fingerprint"`
	} `json:"alerts"`
	GroupLabels struct {
		Alertname string `json:"alertname"`
	} `json:"groupLabels"`
	CommonLabels struct {
		Alertname string `json:"alertname"`
		Severity string `json:"severity"`
	} `json:"commonLabels"`
	CommonAnnotations struct {
		Description string `json:"description"`
		Summary string `json:"summary"`
	} `json:"commonAnnotations"`
	ExternalURL string `json:"externalURL"`
	Version string `json:"version"`
	GroupKey string `json:"groupKey"`
}
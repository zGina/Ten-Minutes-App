package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Technique is
type Technique struct {
	ID                 primitive.ObjectID  `bson:"_id" json:"id"`
	ExternalReferences []ExternalReference `json:"external_references"`
	// ObjectMarkingRefs  []string            `json:"object_marking_refs"`
	// CreatedByRef       string              `json:"created_by_ref"`
	Name        string `json:"name"`
	Description string `json:"description"`
	STIX_ID     string `json:"stix_id"`
	Type        string `json:"type"`
}

type Relationship struct {
	Technique        Technique `json:"-"`
	SourceRef        string    `json:"source_ref"`
	RelationshipType string    `json:"relationship_type"`
	TargetRef        string    `json:"target_ref"`
}

type AttackPattern struct {
	Technique
	KillChainPhases []KillChainPhase
	// XMitreVersion             string   `json:"x_mitre_version"`
	XMitreIsSubtechnique      bool     `bson:"x_mitre_is_subtechnique"  json:"x_mitre_is_subtechnique"`
	XMitrePermissionsRequired []string `json:"x_mitre_permissions_required"`
	XMitreDetection           string   `json:"x_mitre_detection"`
	XMitreDataSources         []string `json:"x_mitre_data_sources"`
	XMitrePlatforms           []string `json:"x_mitre_platforms"`
}

type KillChainPhase struct {
	KillChainName string `json:"kill_chain_name"`
	PhaseName     string `json:"phase_name"`
}

type ExternalReference struct {
	SourceName string `json:"source_name"`
	ExternalID string `json:"external_id,omitempty"`
	URL        string `json:"url"`
}

// KillChainPhase{
// 	KillChainName:"mitre_attack",
// 	PhaseName:"privilege-escalation"
// }

// ExternalReference  {
// 	SourceName :"mitre-attack"
// 	ExternalID :"T1546.004"
// 	URL:"https://attack.mitre.org/techniques/T1546/004"
// }

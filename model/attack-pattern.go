package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExternalReference struct {
	SourceName string `bson:"source_name" json:"source_name"`
	ExternalID string `bson:"external_id" json:"external_id"`
	URL        string `bson:"url" json:"url"`
}

// Technique is
type Technique struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	// ObjectMarkingRefs  []string            `json:"object_marking_refs"`
	// CreatedByRef       string              `bson:"created_by_ref"`
	STIX_ID string `bson:"id" json:"id"`
	Type    string `bson:"type" json:"type"`
}

type Relationship struct {
	Technique        `bson:",inline"`
	SourceRef        string `bson:"source_ref" json:"source_ref"`
	RelationshipType string `bson:"relationship_type" json:"relationship_type"`
	TargetRef        string `bson:"target_ref" json:"target_ref"`
}

type KillChainPhase struct {
	KillChainName string `bson:"kill_chain_name" json:"kill_chain_name"`
	PhaseName     string `bson:"phase_name" json:"phase_name" `
}

type AttackPattern struct {
	Name                      string              `bson:"name" json:"name"`
	ExternalReferences        []ExternalReference `bson:"external_references" json:"external_references"`
	Description               string              `bson:"description" json:"description"`
	Technique                 `bson:",inline"`
	KillChainPhases           []KillChainPhase `bson:"kill_chain_phases" json:"kill_chain_phases"`
	XMitreVersion             string           `bson:"x_mitre_version" json:"x_mitre_version"`
	XMitreIsSubtechnique      bool             `bson:"x_mitre_is_subtechnique" json:"x_mitre_is_subtechnique"`
	XMitreDetection           string           `bson:"x_mitre_detection" json:"x_mitre_detection"`
	XMitrePermissionsRequired []string         `bson:"x_mitre_permissions_required" json:"x_mitre_permissions_required"`
	XMitreDataSources         []string         `bson:"x_mitre_data_sources" json:"x_mitre_data_sources"`
	XMitrePlatforms           []string         `bson:"x_mitre_platforms" json:"x_mitre_platforms"`
}

// New is
func (ap *AttackPattern) New() *AttackPattern {
	return &AttackPattern{
		Technique:                 ap.Technique,
		KillChainPhases:           ap.KillChainPhases,
		XMitreVersion:             ap.XMitreVersion,
		XMitreIsSubtechnique:      ap.XMitreIsSubtechnique,
		XMitreDetection:           ap.XMitreDetection,
		XMitrePermissionsRequired: ap.XMitrePermissionsRequired,
		XMitreDataSources:         ap.XMitreDataSources,
		XMitrePlatforms:           ap.XMitrePlatforms,

		// Created:  time.Now(),
		// Updated:  time.Now(),
	}
}

// New is
func (r *Relationship) New() *Relationship {
	return &Relationship{
		Technique:        r.Technique,
		SourceRef:        r.SourceRef,
		RelationshipType: r.RelationshipType,
		TargetRef:        r.TargetRef,
		// Created:  time.Now(),
		// Updated:  time.Now(),
	}
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

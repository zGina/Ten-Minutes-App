package database

import (
	"github.com/lotteryjs/ten-minutes-app/model"
	"github.com/stretchr/testify/assert"
)

func (s *DatabaseSuite) TestCreateAttackPattern() {
	s.db.DB.Collection("mitre_attack").Drop(nil)

	killChainPhase := model.KillChainPhase{
		KillChainName: "mitre_attack",
		PhaseName:     "privilege-escalation",
	}

	externalReference := model.ExternalReference{
		SourceName: "mitre-attack",
		ExternalID: "T1546.004",
		URL:        "https://attack.mitre.org/techniques/T1546/004",
	}

	technique := model.Technique{
		ExternalReferences: []model.ExternalReference{externalReference},
		Name:               "我是哈哈",
		Description:        "我是嘻嘻",
		STIX_ID:            "id287487",
		Type:               "attack-pattern",
	}

	attackPattern := (&model.AttackPattern{
		Technique:            technique,
		KillChainPhases:      []model.KillChainPhase{killChainPhase},
		XMitreIsSubtechnique: true,
	})

	err := s.db.CreateAttackPattern(attackPattern)
	assert.Nil(s.T(), err)
}

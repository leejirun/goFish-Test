package redfish

import ()

var powerBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Power.Power",
		"@odata.type": "#Power.v1_5_3.Power",
		"@odata.id": "/redfish/v1/Power",
		"Id": "Power-1",
		"Name": "PowerOne",
		"Description": "Power one",
		"PowerControl": [{ 
			"@odata.id": "/redfish/v1/PowerControl",
			"MemberId": "PC1",
			"Name": "Fred", 
			"PhysicalContext": "Upper",
			"PowerAllocatedWatts": 100.0,
			"PowerAvailableWatts": 100.0,
			"PowerCapacityWatts": 100.0,
			"PowerConsumeWatts": 100.0,
			"PowerLimit": {
				"CorrectionInMs": 10000,
				"LimitException": "HardPowerOff",
				"LimitInWatts": 1000.0
			},
			"PowerMetrics": {
				"AverageConsumedWatts": 1000.0,
				"IntervalInMin": 5,
				"MaxConsumedWatts": 1000.0,
				"MinConsumedWatts": 1000.0
			},
			"PowerRequestedWatts": 1000.0,
			"RelatedItem": [],
			"RelatedItem@odata.count": 0,
			"Status": {
				"State": "Enabled",
				"Health": "OK"
			}
		}]
	}`)

var invalidPowerBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Chassis/Members(*)/Self/Power/$entity",
		"@odata.etag": "W/\"1604509181\"",
		"@odata.id": "/redfish/v1/Chassis/Self/Power",
		"@odata.type": "#Power.v1_2_1.Power",
		"Id": "Power",
		"Name": "Power",
		"PowerControl": [
		  {
			"@odata.id": "/redfish/v1/Chassis/Self/Power#/PowerControl/0",
			"MemberId": 0,
			"Name": "Chassis Power Control",
			"PowerLimit": {
			  "CorrectionInMs": 1000,
			  "LimitException": "NoAction",
			  "LimitInWatts": 500
			},
			"PowerMetrics": {
			  "AverageConsumedWatts": 148,
			  "IntervalInMin": 0.083333333333333,
			  "MaxConsumedWatts": 301,
			  "MinConsumedWatts": 0
			},
			"Status": {
			  "Health": "OK",
			  "State": "Enabled"
			}
		  }
		]
	}`)

func TestPower(t *testing.T) {
	var result Power
	err := json.NewDecoder(powerBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Power-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "PowerOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.PowerControl[0].PhysicalContext != common.UpperPhysicalContext {
		t.Errorf("Invalid physical context: %s", result.PowerControl[0].PhysicalContext)
	}

	if result.PowerControl[0].PowerLimit.CorrectionInMs != 10000 {
		t.Errorf("Invalid CorrectionInMs: %d", result.PowerControl[0].PowerLimit.CorrectionInMs)
	}

	if result.PowerControl[0].PowerLimit.LimitException != HardPowerOffPowerLimitException {
		t.Errorf("Invalid LimitException: %s", result.PowerControl[0].PowerLimit.LimitException)
	}

	if result.PowerSupplies[0].IndicatorLED != common.LitIndicatorLED {
		t.Errorf("Invalid PowerSupply IndicatorLED: %s",
			result.PowerSupplies[0].IndicatorLED)
	}

	if result.Voltages[0].MaxReadingRange != 10 {
		t.Errorf("Invalid MaxReadingRange: %f", result.Voltages[0].MaxReadingRange)
	}
}

func TestNonconformingPower(t *testing.T) {
	var result Power
	err := json.NewDecoder(invalidPowerBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.PowerControl[0].MemberID != "0" {
		t.Errorf("Expected first PowerController MemberID to be '0': %s", result.PowerControl[0].MemberID)
	}

	voltage := result.Voltages[0]
	if voltage.MemberID != "218" {
		t.Errorf("Expected first Voltage MemberID to be '218': %s", voltage.MemberID)
	}
}

package equipment

import (
	"atlas-data/xml"
	"github.com/sirupsen/logrus/hooks/test"
	"testing"
)

const testXML = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<imgdir name="01002357.img">
  <imgdir name="info">
    <canvas name="icon" width="36" height="34">
      <vector name="origin" x="1" y="34"/>
    </canvas>
    <canvas name="iconRaw" width="36" height="34">
      <vector name="origin" x="1" y="34"/>
    </canvas>
    <string name="islot" value="Cp"/>
    <string name="vslot" value="CpH1H2H3H4H5HfHsHbAe"/>
    <int name="reqJob" value="0"/>
    <int name="reqLevel" value="50"/>
    <int name="reqSTR" value="0"/>
    <int name="reqDEX" value="0"/>
    <int name="reqINT" value="0"/>
    <int name="reqLUK" value="0"/>
    <int name="incPDD" value="150"/>
    <int name="incMDD" value="150"/>
    <int name="incACC" value="20"/>
    <int name="incEVA" value="20"/>
    <int name="tuc" value="10"/>
    <int name="price" value="500000"/>
    <int name="cash" value="0"/>
    <int name="incSTR" value="15"/>
    <int name="incINT" value="15"/>
    <int name="incDEX" value="15"/>
    <int name="incLUK" value="15"/>
    <int name="only" value="1"/>
    <int name="tradeBlock" value="1"/>
    <int name="tradeAvailable" value="1"/>
    <imgdir name="level">
      <imgdir name="info">
        <imgdir name="1">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="2">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="3">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="4">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="5">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="6">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="7">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="8">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="9">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="10">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="11">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="12">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="13">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="14">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="15">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="16">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="17">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="18">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="19">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="20">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="21">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="22">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="23">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="24">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="25">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="26">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="27">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="28">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="29">
          <int name="exp" value="10000"/>
        </imgdir>
        <imgdir name="30">
          <int name="exp" value="10000"/>
        </imgdir>
      </imgdir>
    </imgdir>
  </imgdir>
  <imgdir name="default">
    <canvas name="default" width="57" height="47">
      <vector name="origin" x="15" y="25"/>
      <imgdir name="map">
        <vector name="brow" x="6" y="-2"/>
      </imgdir>
      <string name="z" value="capBelowAccessory"/>
    </canvas>
  </imgdir>
  <imgdir name="backDefault">
    <canvas name="default" width="49" height="41">
      <vector name="origin" x="15" y="25"/>
      <imgdir name="map">
        <vector name="brow" x="5" y="-5"/>
      </imgdir>
      <string name="z" value="backCap"/>
    </canvas>
  </imgdir>
  <imgdir name="walk1">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="3">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="walk2">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="3">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="stand1">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="stand2">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="alert">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="swingO1">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="swingO2">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="swingO3">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="swingOF">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../backDefault/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="3">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="swingT1">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="swingT2">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="swingT3">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="swingTF">
    <imgdir name="0">
      <uol name="default" value="../../backDefault/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="3">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="swingP1">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="swingP2">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="swingPF">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="3">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="stabO1">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="stabO2">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="stabOF">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="stabT1">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="stabT2">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="stabTF">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="3">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="shoot1">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="shoot2">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="3">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="4">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="shootF">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="proneStab">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="prone">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="heal">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="2">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="fly">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="jump">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="sit">
    <imgdir name="0">
      <uol name="default" value="../../default/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="ladder">
    <imgdir name="0">
      <uol name="default" value="../../backDefault/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../backDefault/default"/>
    </imgdir>
  </imgdir>
  <imgdir name="rope">
    <imgdir name="0">
      <uol name="default" value="../../backDefault/default"/>
    </imgdir>
    <imgdir name="1">
      <uol name="default" value="../../backDefault/default"/>
    </imgdir>
  </imgdir>
</imgdir>
`

func TestReader(t *testing.T) {
	l, _ := test.NewNullLogger()

	rm, err := Read(l)(xml.FromByteArrayProvider([]byte(testXML)))()
	if err != nil {
		t.Fatal(err)
	}
	if rm.Id != 1002357 {
		t.Fatal("id != 1002357")
	}

	if rm.Strength != 15 {
		t.Fatal("strength != 15")
	}
	if rm.Dexterity != 15 {
		t.Fatal("dexterity != 15")
	}
	if rm.Intelligence != 15 {
		t.Fatal("intelligence != 15")
	}
	if rm.Luck != 15 {
		t.Fatal("luck != 15")
	}
	if rm.WeaponDefense != 150 {
		t.Fatal("weapon_defense != 150")
	}
	if rm.MagicDefense != 150 {
		t.Fatal("magic_defense != 150")
	}
	if rm.Accuracy != 20 {
		t.Fatal("accuracy != 20")
	}
	if rm.Avoidability != 20 {
		t.Fatal("avoidability != 20")
	}
	if rm.Slots != 10 {
		t.Fatal("slots != 10")
	}
	if len(rm.EquipSlots) != 1 {
		t.Fatal("len(equip_slots) != 1")
	}
}

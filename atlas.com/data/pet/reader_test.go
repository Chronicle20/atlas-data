package pet

import (
	"atlas-data/xml"
	"github.com/sirupsen/logrus/hooks/test"
	"testing"
)

const testXML = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<imgdir name="5000000.img">
  <imgdir name="info">
    <canvas name="icon" width="32" height="28">
      <vector name="origin" x="0" y="28"/>
    </canvas>
    <canvas name="iconRaw" width="32" height="25">
      <vector name="origin" x="0" y="28"/>
    </canvas>
    <canvas name="iconD" width="29" height="34">
      <vector name="origin" x="-2" y="34"/>
    </canvas>
    <canvas name="iconRawD" width="29" height="34">
      <vector name="origin" x="-2" y="34"/>
    </canvas>
    <int name="hungry" value="2"/>
    <int name="cash" value="1"/>
    <int name="life" value="90"/>
  </imgdir>
  <imgdir name="interact">
    <imgdir name="0">
      <string name="command" value="c1"/>
      <int name="inc" value="1"/>
      <int name="prob" value="30"/>
      <int name="l0" value="1"/>
      <int name="l1" value="9"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="c1_s1"/>
          <string name="1" value="c1_s2"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="stand1"/>
          <string name="0" value="c1_f1"/>
          <string name="1" value="c1_f2"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="1">
      <string name="command" value="c2"/>
      <int name="inc" value="1"/>
      <int name="prob" value="45"/>
      <int name="l0" value="10"/>
      <int name="l1" value="19"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="c2_s1"/>
          <string name="1" value="c2_s2"/>
          <string name="2" value="c2_s3"/>
          <string name="3" value="c2_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="stand1"/>
          <string name="0" value="c2_f1"/>
          <string name="1" value="c2_f2"/>
          <string name="2" value="c2_f3"/>
          <string name="3" value="c2_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="2">
      <string name="command" value="c3"/>
      <int name="inc" value="1"/>
      <int name="prob" value="60"/>
      <int name="l0" value="20"/>
      <int name="l1" value="29"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="c3_s1"/>
          <string name="1" value="c3_s2"/>
          <string name="2" value="c3_s3"/>
          <string name="3" value="c3_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="stand1"/>
          <string name="0" value="c3_f1"/>
          <string name="1" value="c3_f2"/>
          <string name="2" value="c3_f3"/>
          <string name="3" value="c3_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="3">
      <string name="command" value="c4"/>
      <int name="inc" value="1"/>
      <int name="prob" value="75"/>
      <int name="l0" value="30"/>
      <int name="l1" value="30"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="c4_s1"/>
          <string name="1" value="c4_s2"/>
          <string name="2" value="c4_s3"/>
          <string name="3" value="c4_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="prone"/>
          <string name="0" value="c4_f1"/>
          <string name="1" value="c4_f2"/>
          <string name="2" value="c4_f3"/>
          <string name="3" value="c4_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="4">
      <string name="command" value="c5"/>
      <int name="inc" value="1"/>
      <int name="prob" value="30"/>
      <int name="l0" value="1"/>
      <int name="l1" value="9"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="poor"/>
          <string name="0" value="c5_s1"/>
          <string name="1" value="c5_s2"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="angry"/>
          <string name="0" value="c5_f1"/>
          <string name="1" value="c5_f2"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="5">
      <string name="command" value="c6"/>
      <int name="inc" value="1"/>
      <int name="prob" value="45"/>
      <int name="l0" value="10"/>
      <int name="l1" value="19"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="poor"/>
          <string name="0" value="c6_s1"/>
          <string name="1" value="c6_s2"/>
          <string name="2" value="c6_s3"/>
          <string name="3" value="c6_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="angry"/>
          <string name="0" value="c6_f1"/>
          <string name="1" value="c6_f2"/>
          <string name="2" value="c6_f3"/>
          <string name="3" value="c6_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="6">
      <string name="command" value="c7"/>
      <int name="inc" value="1"/>
      <int name="prob" value="60"/>
      <int name="l0" value="20"/>
      <int name="l1" value="29"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="poor"/>
          <string name="0" value="c7_s1"/>
          <string name="1" value="c7_s2"/>
          <string name="2" value="c7_s3"/>
          <string name="3" value="c7_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="angry"/>
          <string name="0" value="c7_f1"/>
          <string name="1" value="c7_f2"/>
          <string name="2" value="c7_f3"/>
          <string name="3" value="c7_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="7">
      <string name="command" value="c8"/>
      <int name="inc" value="1"/>
      <int name="prob" value="75"/>
      <int name="l0" value="30"/>
      <int name="l1" value="30"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="poor"/>
          <string name="0" value="c8_s1"/>
          <string name="1" value="c8_s2"/>
          <string name="2" value="c8_s3"/>
          <string name="3" value="c8_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="stretch"/>
          <string name="0" value="c8_f1"/>
          <string name="1" value="c8_f2"/>
          <string name="2" value="c8_f3"/>
          <string name="3" value="c8_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="8">
      <string name="command" value="c9"/>
      <int name="inc" value="1"/>
      <int name="prob" value="30"/>
      <int name="l0" value="1"/>
      <int name="l1" value="9"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="cry"/>
          <string name="0" value="c9_s1"/>
          <string name="1" value="c9_s2"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="angry"/>
          <string name="0" value="c9_f1"/>
          <string name="1" value="c9_f2"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="9">
      <string name="command" value="c10"/>
      <int name="inc" value="1"/>
      <int name="prob" value="45"/>
      <int name="l0" value="10"/>
      <int name="l1" value="19"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="cry"/>
          <string name="0" value="c10_s1"/>
          <string name="1" value="c10_s2"/>
          <string name="2" value="c10_s3"/>
          <string name="3" value="c10_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="angry"/>
          <string name="0" value="c10_f1"/>
          <string name="1" value="c10_f2"/>
          <string name="2" value="c10_f3"/>
          <string name="3" value="c10_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="10">
      <string name="command" value="c11"/>
      <int name="inc" value="1"/>
      <int name="prob" value="60"/>
      <int name="l0" value="20"/>
      <int name="l1" value="29"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="cry"/>
          <string name="0" value="c11_s1"/>
          <string name="1" value="c11_s2"/>
          <string name="2" value="c11_s3"/>
          <string name="3" value="c11_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="angry"/>
          <string name="0" value="c11_f1"/>
          <string name="1" value="c11_f2"/>
          <string name="2" value="c11_f3"/>
          <string name="4" value="c11_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="11">
      <string name="command" value="c12"/>
      <int name="inc" value="1"/>
      <int name="prob" value="75"/>
      <int name="l0" value="30"/>
      <int name="l1" value="30"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="cry"/>
          <string name="0" value="c12_s1"/>
          <string name="1" value="c12_s2"/>
          <string name="2" value="c12_s3"/>
          <string name="3" value="c12_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="prone"/>
          <string name="0" value="c12_f1"/>
          <string name="1" value="c12_f2"/>
          <string name="2" value="c12_f3"/>
          <string name="3" value="c12_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="12">
      <string name="command" value="c13"/>
      <int name="inc" value="1"/>
      <int name="prob" value="30"/>
      <int name="l0" value="1"/>
      <int name="l1" value="9"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="dung"/>
          <string name="0" value="c13_s1"/>
          <string name="1" value="c13_s2"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="c13_f1"/>
          <string name="1" value="c13_f2"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="13">
      <string name="command" value="c14"/>
      <int name="inc" value="1"/>
      <int name="prob" value="30"/>
      <int name="l0" value="10"/>
      <int name="l1" value="19"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="dung"/>
          <string name="0" value="c14_s1"/>
          <string name="1" value="c14_s2"/>
          <string name="2" value="c14_s3"/>
          <string name="3" value="c14_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="c14_f1"/>
          <string name="1" value="c14_f2"/>
          <string name="2" value="c14_f3"/>
          <string name="3" value="c14_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="14">
      <string name="command" value="c15"/>
      <int name="inc" value="1"/>
      <int name="prob" value="30"/>
      <int name="l0" value="20"/>
      <int name="l1" value="29"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="dung"/>
          <string name="0" value="c15_s1"/>
          <string name="1" value="c15_s2"/>
          <string name="2" value="c15_s3"/>
          <string name="3" value="c15_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="c15_f1"/>
          <string name="1" value="c15_f2"/>
          <string name="2" value="c15_f3"/>
          <string name="3" value="c15_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="15">
      <string name="command" value="c16"/>
      <int name="inc" value="1"/>
      <int name="prob" value="30"/>
      <int name="l0" value="30"/>
      <int name="l1" value="30"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="dung"/>
          <string name="0" value="c16_s1"/>
          <string name="1" value="c16_s2"/>
          <string name="2" value="c16_s3"/>
          <string name="3" value="c16_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="c16_f1"/>
          <string name="1" value="c16_f2"/>
          <string name="2" value="c16_f3"/>
          <string name="3" value="c16_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="16">
      <string name="command" value="c17"/>
      <int name="inc" value="2"/>
      <int name="prob" value="50"/>
      <int name="l0" value="10"/>
      <int name="l1" value="19"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="chat"/>
          <string name="0" value="c17_s1"/>
          <string name="1" value="c17_s2"/>
          <string name="2" value="c17_s3"/>
          <string name="3" value="c17_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="angry"/>
          <string name="0" value="c17_f1"/>
          <string name="1" value="c17_f2"/>
          <string name="2" value="c17_f3"/>
          <string name="3" value="c17_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="17">
      <string name="command" value="c18"/>
      <int name="inc" value="2"/>
      <int name="prob" value="65"/>
      <int name="l0" value="20"/>
      <int name="l1" value="29"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="chat"/>
          <string name="0" value="c18_s1"/>
          <string name="1" value="c18_s2"/>
          <string name="2" value="c18_s3"/>
          <string name="3" value="c18_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="prone"/>
          <string name="0" value="c18_f1"/>
          <string name="1" value="c18_f2"/>
          <string name="2" value="c18_f3"/>
          <string name="3" value="c18_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="18">
      <string name="command" value="c19"/>
      <int name="inc" value="2"/>
      <int name="prob" value="80"/>
      <int name="l0" value="30"/>
      <int name="l1" value="30"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="chat"/>
          <string name="0" value="c19_s1"/>
          <string name="1" value="c19_s2"/>
          <string name="2" value="c19_s3"/>
          <string name="3" value="c19_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="stretch"/>
          <string name="0" value="c19_f1"/>
          <string name="1" value="c19_f2"/>
          <string name="2" value="c19_f3"/>
          <string name="3" value="c19_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="19">
      <string name="command" value="c20"/>
      <int name="inc" value="2"/>
      <int name="prob" value="20"/>
      <int name="l0" value="10"/>
      <int name="l1" value="19"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="alert"/>
          <string name="0" value="c20_s1"/>
          <string name="1" value="c20_s2"/>
          <string name="2" value="c20_s3"/>
          <string name="3" value="c20_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="angry"/>
          <string name="0" value="c20_f1"/>
          <string name="1" value="c20_f2"/>
          <string name="2" value="c20_f3"/>
          <string name="3" value="c20_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="20">
      <string name="command" value="c21"/>
      <int name="inc" value="2"/>
      <int name="prob" value="35"/>
      <int name="l0" value="20"/>
      <int name="l1" value="29"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="alert"/>
          <string name="0" value="c21_s1"/>
          <string name="1" value="c21_s2"/>
          <string name="2" value="c21_s3"/>
          <string name="3" value="c21_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="angry"/>
          <string name="0" value="c21_f1"/>
          <string name="1" value="c21_f2"/>
          <string name="2" value="c21_f3"/>
          <string name="3" value="c21_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="21">
      <string name="command" value="c22"/>
      <int name="inc" value="2"/>
      <int name="prob" value="50"/>
      <int name="l0" value="30"/>
      <int name="l1" value="30"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="alert"/>
          <string name="0" value="c22_s1"/>
          <string name="1" value="c22_s2"/>
          <string name="2" value="c22_s3"/>
          <string name="3" value="c22_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="stretch"/>
          <string name="0" value="c22_f1"/>
          <string name="1" value="c22_f2"/>
          <string name="2" value="c22_f3"/>
          <string name="3" value="c22_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="22">
      <string name="command" value="c23"/>
      <int name="inc" value="3"/>
      <int name="prob" value="30"/>
      <int name="l0" value="20"/>
      <int name="l1" value="29"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="rise"/>
          <string name="0" value="c23_s1"/>
          <string name="1" value="c23_s2"/>
          <string name="2" value="c23_s3"/>
          <string name="3" value="c23_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="angry"/>
          <string name="0" value="c23_f1"/>
          <string name="1" value="c23_f2"/>
          <string name="2" value="c23_f3"/>
          <string name="3" value="c23_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="23">
      <string name="command" value="c24"/>
      <int name="inc" value="3"/>
      <int name="prob" value="40"/>
      <int name="l0" value="30"/>
      <int name="l1" value="30"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="rise"/>
          <string name="0" value="c24_s1"/>
          <string name="1" value="c24_s2"/>
          <string name="2" value="c24_s3"/>
          <string name="3" value="c24_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="stretch"/>
          <string name="0" value="c24_f1"/>
          <string name="1" value="c24_f2"/>
          <string name="2" value="c24_f3"/>
          <string name="3" value="c24_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="24">
      <string name="command" value="c25"/>
      <int name="inc" value="1"/>
      <int name="prob" value="30"/>
      <int name="l0" value="1"/>
      <int name="l1" value="9"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="chat"/>
          <string name="0" value="c25_s1"/>
          <string name="1" value="c25_s2"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="c25_f1"/>
          <string name="1" value="c25_f2"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="25">
      <string name="command" value="c26"/>
      <int name="inc" value="1"/>
      <int name="prob" value="40"/>
      <int name="l0" value="10"/>
      <int name="l1" value="19"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="chat"/>
          <string name="0" value="c26_s1"/>
          <string name="1" value="c26_s2"/>
          <string name="2" value="c26_s3"/>
          <string name="3" value="c26_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="c26_f1"/>
          <string name="1" value="c26_f2"/>
          <string name="2" value="c26_f3"/>
          <string name="3" value="c26_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="26">
      <string name="command" value="c27"/>
      <int name="inc" value="1"/>
      <int name="prob" value="50"/>
      <int name="l0" value="20"/>
      <int name="l1" value="29"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="chat"/>
          <string name="0" value="c27_s1"/>
          <string name="1" value="c27_s2"/>
          <string name="2" value="c27_s3"/>
          <string name="3" value="c27_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="c27_f1"/>
          <string name="1" value="c27_f2"/>
          <string name="2" value="c27_f3"/>
          <string name="3" value="c27_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="27">
      <string name="command" value="c28"/>
      <int name="inc" value="1"/>
      <int name="prob" value="60"/>
      <int name="l0" value="30"/>
      <int name="l1" value="30"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="chat"/>
          <string name="0" value="c28_s1"/>
          <string name="1" value="c28_s2"/>
          <string name="2" value="c28_s3"/>
          <string name="3" value="c28_s4"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="stretch"/>
          <string name="0" value="c28_f1"/>
          <string name="1" value="c28_f2"/>
          <string name="2" value="c28_f3"/>
          <string name="3" value="c28_f4"/>
        </imgdir>
      </imgdir>
    </imgdir>
  </imgdir>
  <imgdir name="food">
    <imgdir name="0">
      <int name="l0" value="1"/>
      <int name="l1" value="9"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="f1_s"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="cry"/>
          <string name="0" value="f1_f"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="1">
      <int name="l0" value="10"/>
      <int name="l1" value="19"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="f2_s"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="cry"/>
          <string name="0" value="f2_f"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="2">
      <int name="l0" value="20"/>
      <int name="l1" value="29"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="f3_s"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="cry"/>
          <string name="0" value="f3_f"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="3">
      <int name="l0" value="30"/>
      <int name="l1" value="30"/>
      <imgdir name="success">
        <imgdir name="0">
          <string name="act" value="rest0"/>
          <string name="0" value="f4_s"/>
        </imgdir>
      </imgdir>
      <imgdir name="fail">
        <imgdir name="0">
          <string name="act" value="cry"/>
          <string name="0" value="f4_f"/>
        </imgdir>
      </imgdir>
    </imgdir>
  </imgdir>
  <imgdir name="slang">
    <imgdir name="0">
      <int name="l0" value="1"/>
      <int name="l1" value="9"/>
      <string name="act" value="rest0"/>
      <string name="0" value="s1"/>
    </imgdir>
    <imgdir name="1">
      <int name="l0" value="10"/>
      <int name="l1" value="19"/>
      <string name="act" value="rest0"/>
      <string name="0" value="s2"/>
    </imgdir>
    <imgdir name="2">
      <int name="l0" value="20"/>
      <int name="l1" value="29"/>
      <string name="act" value="rest0"/>
      <string name="0" value="s3"/>
    </imgdir>
    <imgdir name="3">
      <int name="l0" value="30"/>
      <int name="l1" value="30"/>
      <string name="act" value="rest0"/>
      <string name="0" value="s4"/>
    </imgdir>
  </imgdir>
  <imgdir name="stand0">
    <canvas name="0" width="38" height="39">
      <vector name="origin" x="19" y="39"/>
      <vector name="lt" x="-19" y="-39"/>
      <vector name="rb" x="19" y="0"/>
      <int name="delay" value="1200"/>
    </canvas>
    <canvas name="1" width="38" height="39">
      <vector name="origin" x="19" y="39"/>
      <vector name="lt" x="-19" y="-39"/>
      <vector name="rb" x="19" y="0"/>
      <int name="delay" value="60"/>
    </canvas>
    <canvas name="2" width="38" height="39">
      <vector name="origin" x="19" y="39"/>
      <vector name="lt" x="-19" y="-39"/>
      <vector name="rb" x="19" y="0"/>
      <int name="delay" value="60"/>
    </canvas>
  </imgdir>
  <imgdir name="stand1">
    <canvas name="0" width="38" height="39">
      <vector name="origin" x="19" y="39"/>
      <vector name="lt" x="-19" y="-39"/>
      <vector name="rb" x="19" y="0"/>
      <int name="delay" value="200"/>
    </canvas>
    <canvas name="1" width="36" height="39">
      <vector name="origin" x="19" y="39"/>
      <vector name="lt" x="-19" y="-39"/>
      <vector name="rb" x="17" y="0"/>
      <int name="delay" value="100"/>
    </canvas>
    <canvas name="2" width="35" height="39">
      <vector name="origin" x="19" y="39"/>
      <vector name="lt" x="-19" y="-39"/>
      <vector name="rb" x="16" y="0"/>
      <int name="delay" value="200"/>
    </canvas>
    <uol name="3" value="1"/>
    <uol name="4" value="0"/>
    <canvas name="5" width="41" height="39">
      <vector name="origin" x="19" y="39"/>
      <vector name="lt" x="-19" y="-39"/>
      <vector name="rb" x="22" y="0"/>
      <int name="delay" value="100"/>
    </canvas>
    <canvas name="6" width="44" height="39">
      <vector name="origin" x="19" y="39"/>
      <vector name="lt" x="-19" y="-39"/>
      <vector name="rb" x="25" y="0"/>
      <int name="delay" value="200"/>
    </canvas>
    <canvas name="7" width="48" height="39">
      <vector name="origin" x="19" y="39"/>
      <vector name="lt" x="-19" y="-39"/>
      <vector name="rb" x="29" y="0"/>
      <int name="delay" value="100"/>
    </canvas>
    <canvas name="8" width="45" height="39">
      <vector name="origin" x="19" y="39"/>
      <vector name="lt" x="-19" y="-39"/>
      <vector name="rb" x="26" y="0"/>
      <int name="delay" value="100"/>
    </canvas>
  </imgdir>
  <imgdir name="hungry">
    <canvas name="0" width="49" height="32">
      <vector name="origin" x="17" y="32"/>
      <int name="delay" value="1200"/>
      <vector name="lt" x="-17" y="-32"/>
      <vector name="rb" x="32" y="0"/>
    </canvas>
    <canvas name="1" width="49" height="32">
      <vector name="origin" x="17" y="32"/>
      <int name="delay" value="60"/>
      <vector name="lt" x="-17" y="-32"/>
      <vector name="rb" x="32" y="0"/>
    </canvas>
    <canvas name="2" width="49" height="32">
      <vector name="origin" x="17" y="32"/>
      <int name="delay" value="60"/>
      <vector name="lt" x="-17" y="-32"/>
      <vector name="rb" x="32" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="move">
    <canvas name="0" width="38" height="37">
      <vector name="origin" x="19" y="37"/>
      <int name="delay" value="100"/>
      <vector name="lt" x="-19" y="-37"/>
      <vector name="rb" x="19" y="0"/>
    </canvas>
    <canvas name="1" width="39" height="37">
      <vector name="origin" x="19" y="40"/>
      <int name="delay" value="180"/>
      <vector name="lt" x="-19" y="-37"/>
      <vector name="rb" x="20" y="0"/>
    </canvas>
    <canvas name="2" width="41" height="39">
      <vector name="origin" x="20" y="39"/>
      <int name="delay" value="150"/>
      <vector name="lt" x="-20" y="-39"/>
      <vector name="rb" x="21" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="jump">
    <canvas name="0" width="38" height="35">
      <vector name="origin" x="19" y="35"/>
      <vector name="lt" x="-19" y="-35"/>
      <vector name="rb" x="19" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="rest0">
    <int name="zigzag" value="1"/>
    <canvas name="0" width="46" height="39">
      <vector name="origin" x="20" y="39"/>
      <int name="delay" value="400"/>
      <vector name="lt" x="-20" y="-39"/>
      <vector name="rb" x="26" y="0"/>
    </canvas>
    <canvas name="1" width="44" height="39">
      <vector name="origin" x="20" y="39"/>
      <int name="delay" value="100"/>
      <vector name="lt" x="-20" y="-39"/>
      <vector name="rb" x="24" y="0"/>
    </canvas>
    <canvas name="2" width="41" height="39">
      <vector name="origin" x="20" y="39"/>
      <int name="delay" value="200"/>
      <vector name="lt" x="-20" y="-39"/>
      <vector name="rb" x="21" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="chat">
    <canvas name="0" width="38" height="39">
      <vector name="origin" x="19" y="39"/>
      <int name="delay" value="3000"/>
      <vector name="lt" x="-19" y="-39"/>
      <vector name="rb" x="19" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="angry">
    <canvas name="0" width="37" height="43">
      <vector name="origin" x="21" y="43"/>
      <int name="delay" value="3000"/>
      <vector name="lt" x="-21" y="-43"/>
      <vector name="rb" x="16" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="hang">
    <canvas name="0" width="29" height="56">
      <vector name="origin" x="13" y="60"/>
      <int name="z" value="0"/>
      <vector name="lt" x="-14" y="-28"/>
      <vector name="rb" x="15" y="28"/>
    </canvas>
  </imgdir>
  <imgdir name="rise">
    <canvas name="0" width="41" height="49">
      <vector name="origin" x="20" y="49"/>
      <int name="delay" value="3000"/>
      <vector name="lt" x="-20" y="-49"/>
      <vector name="rb" x="21" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="poor">
    <canvas name="0" width="49" height="32">
      <vector name="origin" x="17" y="32"/>
      <int name="delay" value="3000"/>
      <vector name="lt" x="-17" y="-32"/>
      <vector name="rb" x="32" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="dung">
    <canvas name="0" width="42" height="39">
      <vector name="origin" x="21" y="39"/>
      <int name="delay" value="3000"/>
      <vector name="lt" x="-21" y="-39"/>
      <vector name="rb" x="21" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="stretch">
    <canvas name="0" width="49" height="36">
      <vector name="origin" x="19" y="36"/>
      <int name="delay" value="3000"/>
      <vector name="lt" x="-20" y="-36"/>
      <vector name="rb" x="30" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="prone">
    <canvas name="0" width="54" height="33">
      <vector name="origin" x="22" y="33"/>
      <int name="delay" value="3000"/>
      <vector name="lt" x="-22" y="-33"/>
      <vector name="rb" x="32" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="alert">
    <canvas name="0" width="41" height="34">
      <vector name="origin" x="20" y="34"/>
      <int name="delay" value="3000"/>
      <vector name="lt" x="-20" y="-34"/>
      <vector name="rb" x="21" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="cry">
    <canvas name="0" width="41" height="35">
      <vector name="origin" x="20" y="35"/>
      <int name="delay" value="3000"/>
      <vector name="lt" x="-20" y="-35"/>
      <vector name="rb" x="21" y="0"/>
    </canvas>
  </imgdir>
  <imgdir name="fly">
    <canvas name="0" width="38" height="37">
      <vector name="origin" x="19" y="39"/>
      <int name="delay" value="200"/>
      <vector name="lt" x="-19" y="-37"/>
      <vector name="rb" x="19" y="0"/>
      <int name="z" value="1"/>
    </canvas>
    <canvas name="1" width="41" height="39">
      <vector name="origin" x="20" y="39"/>
      <int name="delay" value="200"/>
      <vector name="lt" x="-19" y="-37"/>
      <vector name="rb" x="20" y="0"/>
    </canvas>
  </imgdir>
</imgdir>
`

func TestReader(t *testing.T) {
	l, _ := test.NewNullLogger()

	rm, err := Read(l)(xml.FromByteArrayProvider([]byte(testXML)))()
	if err != nil {
		t.Fatal(err)
	}
	if rm.Id != 5000000 {
		t.Fatal("id != 5000000")
	}
	if rm.Hungry != 2 {
		t.Fatal("hungry != 2")
	}
	if !rm.Cash {
		t.Fatal("cash != true")
	}
	if rm.Life != 90 {
		t.Fatal("life != 90")
	}
	if len(rm.Skills) != 28 {
		t.Fatal("len(rm.Skills) != 28")
	}
}

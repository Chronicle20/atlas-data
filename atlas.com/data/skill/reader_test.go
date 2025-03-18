package skill

import (
	"atlas-data/skill/effect"
	"atlas-data/xml"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/sirupsen/logrus/hooks/test"
	"testing"
)

const testXML = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<imgdir name="000.img">
  <imgdir name="info">
    <canvas name="icon" width="26" height="30">
      <vector name="origin" x="-4" y="30"/>
    </canvas>
  </imgdir>
  <imgdir name="skill">
    <imgdir name="0001003">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <int name="invisible" value="1"/>
      <imgdir name="level">
        <imgdir name="1">
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="0001004">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <int name="x" value="1"/>
          <int name="time" value="2100000"/>
        </imgdir>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="52" height="14">
          <vector name="origin" x="27" y="12"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="1" width="58" height="18">
          <vector name="origin" x="30" y="15"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="60" height="21">
          <vector name="origin" x="31" y="18"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="60" height="60">
          <vector name="origin" x="31" y="57"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="71" height="118">
          <vector name="origin" x="38" y="115"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="144" height="115">
          <vector name="origin" x="76" y="109"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="147" height="119">
          <vector name="origin" x="77" y="111"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="150" height="122">
          <vector name="origin" x="78" y="112"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="150" height="124">
          <vector name="origin" x="76" y="113"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="140" height="99">
          <vector name="origin" x="75" y="96"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="142" height="99">
          <vector name="origin" x="75" y="93"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="147" height="101">
          <vector name="origin" x="77" y="90"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="12" width="148" height="100">
          <vector name="origin" x="76" y="85"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="13" width="146" height="94">
          <vector name="origin" x="73" y="77"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="14" width="130" height="62">
          <vector name="origin" x="74" y="63"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="15" width="35" height="37">
          <vector name="origin" x="-22" y="40"/>
          <int name="delay" value="60"/>
        </canvas>
      </imgdir>
      <imgdir name="effect0">
        <canvas name="0" width="1" height="1">
          <vector name="origin" x="0" y="0"/>
          <int name="delay" value="300"/>
        </canvas>
        <canvas name="1" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="152" height="167">
          <vector name="origin" x="82" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="151" height="168">
          <vector name="origin" x="79" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="165" height="168">
          <vector name="origin" x="76" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="163" height="168">
          <vector name="origin" x="72" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="181" height="167">
          <vector name="origin" x="93" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="178" height="166">
          <vector name="origin" x="93" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="174" height="167">
          <vector name="origin" x="91" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <int name="z" value="-2"/>
      </imgdir>
      <int name="invisible" value="1"/>
    </imgdir>
    <imgdir name="0001005">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="mpCon" value="30"/>
          <int name="x" value="4"/>
          <int name="time" value="2400"/>
          <vector name="lt" x="-400" y="-300"/>
          <vector name="rb" x="400" y="300"/>
          <int name="cooltime" value="7200"/>
        </imgdir>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="150" height="124">
          <vector name="origin" x="74" y="90"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="1" width="133" height="155">
          <vector name="origin" x="62" y="101"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="2" width="123" height="161">
          <vector name="origin" x="59" y="114"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="3" width="106" height="154">
          <vector name="origin" x="52" y="117"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="4" width="86" height="144">
          <vector name="origin" x="43" y="119"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="5" width="65" height="143">
          <vector name="origin" x="34" y="121"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="6" width="71" height="144">
          <vector name="origin" x="39" y="123"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="7" width="177" height="194">
          <vector name="origin" x="92" y="150"/>
          <int name="delay" value="150"/>
        </canvas>
        <canvas name="8" width="173" height="194">
          <vector name="origin" x="89" y="153"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="9" width="164" height="178">
          <vector name="origin" x="85" y="141"/>
          <int name="delay" value="150"/>
        </canvas>
        <canvas name="10" width="151" height="158">
          <vector name="origin" x="74" y="125"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="11" width="53" height="125">
          <vector name="origin" x="30" y="119"/>
          <int name="delay" value="120"/>
        </canvas>
      </imgdir>
      <imgdir name="affected">
        <canvas name="0" width="4" height="4">
          <vector name="origin" x="0" y="4"/>
          <int name="delay" value="630"/>
        </canvas>
        <canvas name="1" width="104" height="114">
          <vector name="origin" x="54" y="114"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="2" width="99" height="128">
          <vector name="origin" x="52" y="128"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="3" width="93" height="133">
          <vector name="origin" x="49" y="133"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="4" width="96" height="136">
          <vector name="origin" x="47" y="136"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="5" width="104" height="141">
          <vector name="origin" x="45" y="141"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="6" width="104" height="142">
          <vector name="origin" x="45" y="142"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="7" width="123" height="147">
          <vector name="origin" x="61" y="147"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="8" width="113" height="144">
          <vector name="origin" x="58" y="144"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="9" width="109" height="140">
          <vector name="origin" x="56" y="140"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="10" width="104" height="139">
          <vector name="origin" x="54" y="137"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="11" width="99" height="145">
          <vector name="origin" x="52" y="133"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="12" width="93" height="137">
          <vector name="origin" x="49" y="126"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="13" width="82" height="129">
          <vector name="origin" x="41" y="122"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="14" width="70" height="89">
          <vector name="origin" x="36" y="82"/>
          <int name="delay" value="90"/>
        </canvas>
      </imgdir>
      <int name="invisible" value="1"/>
    </imgdir>
    <imgdir name="0000008">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <int name="invisible" value="1"/>
      <imgdir name="level">
        <imgdir name="1">
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="0001006">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
        </imgdir>
      </imgdir>
      <int name="invisible" value="1"/>
    </imgdir>
    <imgdir name="0001001">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="mpCon" value="5"/>
          <int name="time" value="30"/>
          <int name="x" value="4"/>
          <int name="cooltime" value="120"/>
        </imgdir>
        <imgdir name="2">
          <string name="hs" value="h2"/>
          <int name="mpCon" value="10"/>
          <int name="time" value="30"/>
          <int name="x" value="8"/>
          <int name="cooltime" value="120"/>
        </imgdir>
        <imgdir name="3">
          <string name="hs" value="h3"/>
          <int name="mpCon" value="15"/>
          <int name="time" value="30"/>
          <int name="x" value="12"/>
          <int name="cooltime" value="120"/>
        </imgdir>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="47" height="77">
          <vector name="origin" x="27" y="77"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="1" width="49" height="81">
          <vector name="origin" x="26" y="81"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="2" width="53" height="85">
          <vector name="origin" x="29" y="85"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="3" width="55" height="90">
          <vector name="origin" x="30" y="89"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="4" width="57" height="94">
          <vector name="origin" x="32" y="93"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="5" width="56" height="98">
          <vector name="origin" x="33" y="97"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="6" width="52" height="102">
          <vector name="origin" x="31" y="101"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="7" width="55" height="106">
          <vector name="origin" x="31" y="105"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="8" width="52" height="110">
          <vector name="origin" x="30" y="109"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="9" width="51" height="115">
          <vector name="origin" x="27" y="114"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="10" width="53" height="118">
          <vector name="origin" x="29" y="118"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="11" width="53" height="104">
          <vector name="origin" x="29" y="104"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="12" width="53" height="64">
          <vector name="origin" x="30" y="65"/>
          <int name="delay" value="90"/>
        </canvas>
      </imgdir>
    </imgdir>
    <imgdir name="0001002">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="mpCon" value="4"/>
          <int name="time" value="4"/>
          <int name="speed" value="10"/>
          <int name="cooltime" value="60"/>
        </imgdir>
        <imgdir name="2">
          <string name="hs" value="h2"/>
          <int name="mpCon" value="7"/>
          <int name="time" value="8"/>
          <int name="speed" value="15"/>
          <int name="cooltime" value="60"/>
        </imgdir>
        <imgdir name="3">
          <string name="hs" value="h3"/>
          <int name="mpCon" value="10"/>
          <int name="time" value="12"/>
          <int name="speed" value="20"/>
          <int name="cooltime" value="60"/>
        </imgdir>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="59" height="53">
          <vector name="origin" x="33" y="53"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="1" width="61" height="59">
          <vector name="origin" x="33" y="57"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="2" width="66" height="65">
          <vector name="origin" x="34" y="59"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="3" width="70" height="71">
          <vector name="origin" x="35" y="65"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="4" width="70" height="75">
          <vector name="origin" x="36" y="69"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="5" width="72" height="81">
          <vector name="origin" x="35" y="74"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="6" width="70" height="91">
          <vector name="origin" x="33" y="87"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="7" width="68" height="106">
          <vector name="origin" x="33" y="103"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="8" width="64" height="97">
          <vector name="origin" x="34" y="94"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="9" width="60" height="116">
          <vector name="origin" x="34" y="113"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="10" width="58" height="62">
          <vector name="origin" x="33" y="60"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="11" width="57" height="58">
          <vector name="origin" x="32" y="57"/>
          <int name="delay" value="90"/>
        </canvas>
      </imgdir>
    </imgdir>
    <imgdir name="0001000">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="mpCon" value="3"/>
          <int name="fixdamage" value="10"/>
          <imgdir name="ball">
            <canvas name="0" width="32" height="35">
              <vector name="origin" x="17" y="17"/>
              <int name="delay" value="90"/>
            </canvas>
            <canvas name="1" width="33" height="31">
              <vector name="origin" x="19" y="13"/>
              <int name="delay" value="90"/>
            </canvas>
            <canvas name="2" width="34" height="30">
              <vector name="origin" x="19" y="16"/>
              <int name="delay" value="90"/>
            </canvas>
          </imgdir>
          <imgdir name="hit">
            <imgdir name="0">
              <canvas name="0" width="76" height="79">
                <vector name="origin" x="37" y="70"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="1" width="71" height="75">
                <vector name="origin" x="33" y="68"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="2" width="70" height="71">
                <vector name="origin" x="33" y="67"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="3" width="70" height="68">
                <vector name="origin" x="34" y="65"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="4" width="71" height="61">
                <vector name="origin" x="39" y="60"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="5" width="73" height="62">
                <vector name="origin" x="42" y="61"/>
                <int name="delay" value="90"/>
              </canvas>
            </imgdir>
          </imgdir>
        </imgdir>
        <imgdir name="2">
          <string name="hs" value="h2"/>
          <int name="mpCon" value="5"/>
          <int name="fixdamage" value="25"/>
          <imgdir name="ball">
            <canvas name="0" width="32" height="35">
              <vector name="origin" x="17" y="17"/>
              <int name="delay" value="90"/>
            </canvas>
            <canvas name="1" width="33" height="31">
              <vector name="origin" x="19" y="13"/>
              <int name="delay" value="90"/>
            </canvas>
            <canvas name="2" width="34" height="30">
              <vector name="origin" x="19" y="16"/>
              <int name="delay" value="90"/>
            </canvas>
          </imgdir>
          <imgdir name="hit">
            <imgdir name="0">
              <canvas name="0" width="76" height="79">
                <vector name="origin" x="37" y="70"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="1" width="71" height="75">
                <vector name="origin" x="33" y="68"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="2" width="70" height="71">
                <vector name="origin" x="33" y="67"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="3" width="70" height="68">
                <vector name="origin" x="34" y="65"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="4" width="71" height="61">
                <vector name="origin" x="39" y="60"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="5" width="73" height="62">
                <vector name="origin" x="42" y="61"/>
                <int name="delay" value="90"/>
              </canvas>
            </imgdir>
          </imgdir>
        </imgdir>
        <imgdir name="3">
          <string name="hs" value="h3"/>
          <int name="mpCon" value="7"/>
          <int name="fixdamage" value="40"/>
          <imgdir name="ball">
            <canvas name="0" width="32" height="35">
              <vector name="origin" x="17" y="17"/>
              <int name="delay" value="90"/>
            </canvas>
            <canvas name="1" width="33" height="31">
              <vector name="origin" x="19" y="13"/>
              <int name="delay" value="90"/>
            </canvas>
            <canvas name="2" width="34" height="30">
              <vector name="origin" x="19" y="16"/>
              <int name="delay" value="90"/>
            </canvas>
          </imgdir>
          <imgdir name="hit">
            <imgdir name="0">
              <canvas name="0" width="76" height="79">
                <vector name="origin" x="37" y="70"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="1" width="71" height="75">
                <vector name="origin" x="33" y="68"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="2" width="70" height="71">
                <vector name="origin" x="33" y="67"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="3" width="70" height="68">
                <vector name="origin" x="34" y="65"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="4" width="71" height="61">
                <vector name="origin" x="39" y="60"/>
                <int name="delay" value="90"/>
              </canvas>
              <canvas name="5" width="73" height="62">
                <vector name="origin" x="42" y="61"/>
                <int name="delay" value="90"/>
              </canvas>
            </imgdir>
          </imgdir>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="0001007">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <int name="invisible" value="1"/>
      <int name="disable" value="1"/>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
        </imgdir>
        <imgdir name="2">
          <string name="hs" value="h2"/>
        </imgdir>
        <imgdir name="3">
          <string name="hs" value="h3"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="0001009">
      <string name="info" value="일격필살"/>
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="action">
        <string name="0" value="bamboo"/>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="1" height="1">
          <vector name="origin" x="0" y="0"/>
          <int name="z" value="0"/>
          <int name="delay" value="1080"/>
        </canvas>
        <canvas name="1" width="65" height="72">
          <vector name="origin" x="98" y="335"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="2" width="65" height="72">
          <vector name="origin" x="98" y="349"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="65" height="72">
          <vector name="origin" x="98" y="356"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="4" width="66" height="100">
          <vector name="origin" x="98" y="389"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="5" width="87" height="94">
          <vector name="origin" x="120" y="387"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="6" width="86" height="73">
          <vector name="origin" x="119" y="369"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="7" width="89" height="93">
          <vector name="origin" x="120" y="389"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="8" width="108" height="94">
          <vector name="origin" x="118" y="385"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="9" width="97" height="88">
          <vector name="origin" x="109" y="382"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="10" width="96" height="94">
          <vector name="origin" x="107" y="388"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="11" width="111" height="101">
          <vector name="origin" x="123" y="387"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="12" width="98" height="97">
          <vector name="origin" x="121" y="385"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="13" width="84" height="92">
          <vector name="origin" x="108" y="385"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="14" width="451" height="126">
          <vector name="origin" x="307" y="405"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="15" width="499" height="176">
          <vector name="origin" x="325" y="429"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="16" width="606" height="212">
          <vector name="origin" x="380" y="447"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="17" width="645" height="225">
          <vector name="origin" x="399" y="452"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="18" width="669" height="230">
          <vector name="origin" x="413" y="458"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="27" width="1" height="1">
          <vector name="origin" x="0" y="0"/>
          <int name="z" value="0"/>
          <int name="delay" value="1200"/>
        </canvas>
      </imgdir>
      <imgdir name="hit">
        <imgdir name="0">
          <canvas name="0" width="124" height="109">
            <vector name="origin" x="61" y="94"/>
            <int name="z" value="0"/>
          </canvas>
          <canvas name="1" width="96" height="98">
            <vector name="origin" x="46" y="87"/>
            <int name="z" value="0"/>
          </canvas>
          <canvas name="2" width="94" height="92">
            <vector name="origin" x="47" y="83"/>
            <int name="z" value="0"/>
          </canvas>
          <canvas name="3" width="99" height="83">
            <vector name="origin" x="57" y="77"/>
            <int name="z" value="0"/>
          </canvas>
          <canvas name="4" width="106" height="89">
            <vector name="origin" x="59" y="81"/>
            <int name="z" value="0"/>
          </canvas>
        </imgdir>
      </imgdir>
      <imgdir name="level">
        <imgdir name="1">
          <int name="damagepc" value="100"/>
          <int name="mobCount" value="30"/>
          <vector name="lt" x="-1000" y="-300"/>
          <vector name="rb" x="1000" y="15"/>
        </imgdir>
      </imgdir>
      <int name="invisible" value="1"/>
      <imgdir name="screen">
        <canvas name="0" width="1" height="1">
          <vector name="origin" x="0" y="0"/>
          <int name="z" value="0"/>
          <int name="delay" value="2220"/>
        </canvas>
        <canvas name="1" width="708" height="97">
          <vector name="origin" x="375" y="190"/>
          <int name="z" value="0"/>
          <int name="delay" value="240"/>
        </canvas>
        <canvas name="2" width="721" height="97">
          <vector name="origin" x="383" y="190"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="732" height="97">
          <vector name="origin" x="388" y="190"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="4" width="742" height="97">
          <vector name="origin" x="393" y="190"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="5" width="744" height="178">
          <vector name="origin" x="394" y="179"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="6" width="780" height="293">
          <vector name="origin" x="398" y="94"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="7" width="790" height="281">
          <vector name="origin" x="403" y="83"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="8" width="791" height="276">
          <vector name="origin" x="402" y="78"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="9" width="798" height="273">
          <vector name="origin" x="403" y="75"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="10" width="798" height="271">
          <vector name="origin" x="402" y="74"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="11" width="798" height="269">
          <vector name="origin" x="407" y="73"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="12" width="756" height="266">
          <vector name="origin" x="396" y="72"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
      </imgdir>
    </imgdir>
    <imgdir name="0001010">
      <string name="info" value="무적"/>
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="effect">
        <canvas name="0" width="74" height="15">
          <vector name="origin" x="38" y="13"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="1" width="88" height="17">
          <vector name="origin" x="45" y="15"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="2" width="112" height="62">
          <vector name="origin" x="58" y="61"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="3" width="123" height="87">
          <vector name="origin" x="63" y="85"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="4" width="154" height="172">
          <vector name="origin" x="79" y="171"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="5" width="153" height="170">
          <vector name="origin" x="77" y="168"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="6" width="150" height="168">
          <vector name="origin" x="76" y="166"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="7" width="140" height="163">
          <vector name="origin" x="71" y="162"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="8" width="120" height="138">
          <vector name="origin" x="61" y="153"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="9" width="106" height="140">
          <vector name="origin" x="60" y="140"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="10" width="89" height="70">
          <vector name="origin" x="56" y="131"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="11" width="67" height="53">
          <vector name="origin" x="41" y="131"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="12" width="152" height="147">
          <vector name="origin" x="76" y="176"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="13" width="158" height="141">
          <vector name="origin" x="78" y="174"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="14" width="163" height="140">
          <vector name="origin" x="79" y="167"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="15" width="121" height="141">
          <vector name="origin" x="61" y="161"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="16" width="116" height="143">
          <vector name="origin" x="58" y="159"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="17" width="92" height="102">
          <vector name="origin" x="49" y="108"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="18" width="90" height="98">
          <vector name="origin" x="48" y="104"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="19" width="86" height="102">
          <vector name="origin" x="46" y="110"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="20" width="66" height="109">
          <vector name="origin" x="36" y="120"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="21" width="66" height="112">
          <vector name="origin" x="36" y="123"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="22" width="54" height="53">
          <vector name="origin" x="36" y="123"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="23" width="49" height="46">
          <vector name="origin" x="31" y="117"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="24" width="43" height="47">
          <vector name="origin" x="25" y="117"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="25" width="44" height="49">
          <vector name="origin" x="26" y="118"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="26" width="38" height="34">
          <vector name="origin" x="20" y="104"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
      </imgdir>
      <imgdir name="affected">
        <int name="pos" value="1"/>
        <int name="repeat" value="1"/>
        <canvas name="0" width="38" height="34">
          <vector name="origin" x="20" y="103"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="1" width="38" height="34">
          <vector name="origin" x="20" y="104"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="2" width="38" height="34">
          <vector name="origin" x="20" y="105"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="3" width="38" height="34">
          <vector name="origin" x="20" y="106"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="4" width="38" height="34">
          <vector name="origin" x="20" y="105"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="5" width="38" height="34">
          <vector name="origin" x="20" y="104"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="6" width="38" height="34">
          <vector name="origin" x="20" y="103"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="7" width="38" height="34">
          <vector name="origin" x="20" y="102"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
      </imgdir>
      <imgdir name="action">
        <string name="0" value="alert2"/>
      </imgdir>
      <imgdir name="level">
        <imgdir name="1">
          <int name="time" value="30"/>
          <int name="x" value="1"/>
        </imgdir>
      </imgdir>
      <int name="invisible" value="1"/>
      <imgdir name="repeat">
        <canvas name="0" width="106" height="106">
          <vector name="origin" x="52" y="90"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="1" width="106" height="106">
          <vector name="origin" x="52" y="90"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="2" width="108" height="108">
          <vector name="origin" x="53" y="91"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="108" height="108">
          <vector name="origin" x="53" y="91"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="4" width="110" height="110">
          <vector name="origin" x="54" y="92"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="5" width="110" height="110">
          <vector name="origin" x="54" y="92"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="6" width="110" height="110">
          <vector name="origin" x="54" y="92"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="7" width="108" height="108">
          <vector name="origin" x="53" y="91"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="8" width="108" height="108">
          <vector name="origin" x="53" y="91"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="9" width="106" height="106">
          <vector name="origin" x="52" y="90"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
      </imgdir>
    </imgdir>
    <imgdir name="0001011">
      <string name="info" value="광폭"/>
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="action">
        <string name="0" value="alert2"/>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="74" height="15">
          <vector name="origin" x="39" y="13"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="1" width="88" height="17">
          <vector name="origin" x="46" y="15"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="2" width="112" height="62">
          <vector name="origin" x="59" y="60"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="3" width="123" height="87">
          <vector name="origin" x="64" y="85"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="4" width="154" height="172">
          <vector name="origin" x="80" y="171"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="5" width="153" height="170">
          <vector name="origin" x="78" y="168"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="6" width="150" height="168">
          <vector name="origin" x="77" y="166"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="7" width="140" height="163">
          <vector name="origin" x="72" y="162"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="8" width="120" height="138">
          <vector name="origin" x="62" y="153"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="9" width="106" height="140">
          <vector name="origin" x="61" y="140"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="10" width="89" height="70">
          <vector name="origin" x="57" y="131"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="11" width="67" height="53">
          <vector name="origin" x="42" y="131"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="12" width="152" height="147">
          <vector name="origin" x="77" y="176"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="13" width="158" height="141">
          <vector name="origin" x="79" y="174"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="14" width="163" height="138">
          <vector name="origin" x="80" y="167"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="15" width="121" height="141">
          <vector name="origin" x="62" y="161"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="16" width="116" height="143">
          <vector name="origin" x="59" y="159"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="17" width="92" height="106">
          <vector name="origin" x="50" y="112"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="18" width="90" height="102">
          <vector name="origin" x="49" y="108"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="19" width="86" height="102">
          <vector name="origin" x="47" y="110"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="20" width="66" height="109">
          <vector name="origin" x="37" y="120"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="21" width="66" height="112">
          <vector name="origin" x="37" y="123"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="22" width="51" height="55">
          <vector name="origin" x="37" y="123"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="23" width="45" height="48">
          <vector name="origin" x="32" y="117"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="24" width="39" height="49">
          <vector name="origin" x="26" y="117"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="25" width="40" height="51">
          <vector name="origin" x="27" y="118"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="26" width="30" height="40">
          <vector name="origin" x="17" y="108"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
      </imgdir>
      <imgdir name="affected">
        <canvas name="0" width="30" height="40">
          <vector name="origin" x="20" y="60"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="1" width="30" height="40">
          <vector name="origin" x="20" y="61"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="2" width="30" height="40">
          <vector name="origin" x="20" y="62"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="3" width="30" height="40">
          <vector name="origin" x="20" y="63"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="4" width="30" height="40">
          <vector name="origin" x="20" y="64"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="5" width="30" height="40">
          <vector name="origin" x="20" y="63"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="6" width="30" height="40">
          <vector name="origin" x="20" y="62"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <canvas name="7" width="30" height="40">
          <vector name="origin" x="20" y="61"/>
          <int name="z" value="0"/>
          <int name="delay" value="90"/>
        </canvas>
        <int name="pos" value="1"/>
        <int name="repeat" value="1"/>
      </imgdir>
      <imgdir name="level">
        <imgdir name="1">
          <int name="time" value="30"/>
          <int name="x" value="-3"/>
        </imgdir>
      </imgdir>
      <int name="invisible" value="1"/>
      <imgdir name="repeat">
        <canvas name="0" width="84" height="98">
          <vector name="origin" x="38" y="87"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="1" width="84" height="109">
          <vector name="origin" x="39" y="97"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="2" width="84" height="112">
          <vector name="origin" x="41" y="99"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="86" height="112">
          <vector name="origin" x="43" y="99"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="4" width="92" height="112">
          <vector name="origin" x="46" y="99"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="5" width="94" height="111">
          <vector name="origin" x="48" y="99"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="6" width="93" height="112">
          <vector name="origin" x="49" y="102"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="7" width="96" height="109">
          <vector name="origin" x="53" y="99"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="8" width="92" height="112">
          <vector name="origin" x="51" y="101"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="9" width="88" height="113">
          <vector name="origin" x="47" y="102"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="10" width="85" height="103">
          <vector name="origin" x="45" y="92"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="11" width="80" height="103">
          <vector name="origin" x="41" y="92"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="12" width="80" height="106">
          <vector name="origin" x="40" y="96"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="13" width="83" height="108">
          <vector name="origin" x="39" y="97"/>
          <int name="delay" value="120"/>
        </canvas>
        <int name="z" value="-8"/>
      </imgdir>
    </imgdir>
    <imgdir name="0000012">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="x" value="1"/>
          <int name="y" value="2"/>
          <int name="z" value="1"/>
        </imgdir>
        <imgdir name="2">
          <string name="hs" value="h2"/>
          <int name="x" value="2"/>
          <int name="y" value="4"/>
          <int name="z" value="2"/>
        </imgdir>
        <imgdir name="3">
          <string name="hs" value="h3"/>
          <int name="x" value="3"/>
          <int name="y" value="6"/>
          <int name="z" value="3"/>
        </imgdir>
        <imgdir name="4">
          <string name="hs" value="h4"/>
          <int name="x" value="4"/>
          <int name="y" value="8"/>
          <int name="z" value="4"/>
        </imgdir>
        <imgdir name="5">
          <string name="hs" value="h5"/>
          <int name="x" value="5"/>
          <int name="y" value="10"/>
          <int name="z" value="5"/>
        </imgdir>
        <imgdir name="6">
          <string name="hs" value="h6"/>
          <int name="x" value="6"/>
          <int name="y" value="12"/>
          <int name="z" value="6"/>
        </imgdir>
        <imgdir name="7">
          <string name="hs" value="h7"/>
          <int name="x" value="7"/>
          <int name="y" value="14"/>
          <int name="z" value="7"/>
        </imgdir>
        <imgdir name="8">
          <string name="hs" value="h8"/>
          <int name="x" value="8"/>
          <int name="y" value="16"/>
          <int name="z" value="8"/>
        </imgdir>
        <imgdir name="9">
          <string name="hs" value="h9"/>
          <int name="x" value="9"/>
          <int name="y" value="18"/>
          <int name="z" value="9"/>
        </imgdir>
        <imgdir name="10">
          <string name="hs" value="h10"/>
          <int name="x" value="10"/>
          <int name="y" value="20"/>
          <int name="z" value="10"/>
        </imgdir>
        <imgdir name="11">
          <string name="hs" value="h11"/>
          <int name="x" value="11"/>
          <int name="y" value="22"/>
          <int name="z" value="11"/>
        </imgdir>
        <imgdir name="12">
          <string name="hs" value="h12"/>
          <int name="x" value="12"/>
          <int name="y" value="24"/>
          <int name="z" value="12"/>
        </imgdir>
        <imgdir name="13">
          <string name="hs" value="h13"/>
          <int name="x" value="13"/>
          <int name="y" value="26"/>
          <int name="z" value="13"/>
        </imgdir>
        <imgdir name="14">
          <string name="hs" value="h14"/>
          <int name="x" value="14"/>
          <int name="y" value="28"/>
          <int name="z" value="14"/>
        </imgdir>
        <imgdir name="15">
          <string name="hs" value="h15"/>
          <int name="x" value="15"/>
          <int name="y" value="30"/>
          <int name="z" value="15"/>
        </imgdir>
        <imgdir name="16">
          <string name="hs" value="h16"/>
          <int name="x" value="16"/>
          <int name="y" value="32"/>
          <int name="z" value="16"/>
        </imgdir>
        <imgdir name="17">
          <string name="hs" value="h17"/>
          <int name="x" value="17"/>
          <int name="y" value="34"/>
          <int name="z" value="17"/>
        </imgdir>
        <imgdir name="18">
          <string name="hs" value="h18"/>
          <int name="x" value="18"/>
          <int name="y" value="36"/>
          <int name="z" value="18"/>
        </imgdir>
        <imgdir name="19">
          <string name="hs" value="h19"/>
          <int name="x" value="19"/>
          <int name="y" value="38"/>
          <int name="z" value="19"/>
        </imgdir>
        <imgdir name="20">
          <string name="hs" value="h20"/>
          <int name="x" value="20"/>
          <int name="y" value="40"/>
          <int name="z" value="20"/>
        </imgdir>
      </imgdir>
      <int name="disable" value="1"/>
      <int name="invisible" value="1"/>
    </imgdir>
    <imgdir name="0001013">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="mpCon" value="10"/>
          <int name="time" value="2100000"/>
          <int name="pdd" value="10"/>
          <int name="mdd" value="10"/>
          <int name="x" value="1"/>
          <string name="dateExpire" value="2009060800"/>
        </imgdir>
        <imgdir name="2">
          <string name="hs" value="h2"/>
          <int name="mpCon" value="10"/>
          <int name="time" value="2100000"/>
          <int name="pdd" value="20"/>
          <int name="mdd" value="20"/>
          <int name="x" value="1"/>
          <string name="dateExpire" value="2009060800"/>
        </imgdir>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="52" height="14">
          <vector name="origin" x="27" y="12"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="1" width="58" height="18">
          <vector name="origin" x="30" y="15"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="60" height="21">
          <vector name="origin" x="31" y="18"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="60" height="60">
          <vector name="origin" x="31" y="57"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="71" height="118">
          <vector name="origin" x="38" y="115"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="144" height="115">
          <vector name="origin" x="76" y="109"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="147" height="119">
          <vector name="origin" x="77" y="111"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="150" height="122">
          <vector name="origin" x="78" y="112"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="150" height="124">
          <vector name="origin" x="76" y="113"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="140" height="99">
          <vector name="origin" x="75" y="96"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="142" height="99">
          <vector name="origin" x="75" y="93"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="147" height="101">
          <vector name="origin" x="77" y="90"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="12" width="148" height="100">
          <vector name="origin" x="76" y="85"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="13" width="146" height="94">
          <vector name="origin" x="73" y="77"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="14" width="130" height="62">
          <vector name="origin" x="74" y="63"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="15" width="35" height="37">
          <vector name="origin" x="-22" y="40"/>
          <int name="delay" value="60"/>
        </canvas>
      </imgdir>
      <imgdir name="effect0">
        <canvas name="0" width="1" height="1">
          <vector name="origin" x="0" y="0"/>
          <int name="delay" value="300"/>
        </canvas>
        <canvas name="1" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="152" height="167">
          <vector name="origin" x="82" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="151" height="168">
          <vector name="origin" x="79" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="165" height="168">
          <vector name="origin" x="76" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="163" height="168">
          <vector name="origin" x="72" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="181" height="167">
          <vector name="origin" x="93" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="178" height="166">
          <vector name="origin" x="93" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="174" height="167">
          <vector name="origin" x="91" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <int name="z" value="-2"/>
      </imgdir>
      <int name="invisible" value="1"/>
      <int name="timeLimited" value="1"/>
      <int name="disable" value="1"/>
    </imgdir>
    <imgdir name="0001014">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="mpCon" value="10"/>
          <int name="x" value="50"/>
          <int name="y" value="1"/>
          <int name="time" value="1"/>
          <string name="dateExpire" value="2009060800"/>
        </imgdir>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="89" height="53">
          <vector name="origin" x="7" y="47"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="1" width="57" height="46">
          <vector name="origin" x="-3" y="41"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="2" width="53" height="44">
          <vector name="origin" x="-5" y="40"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="63" height="32">
          <vector name="origin" x="-8" y="30"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="4" width="52" height="22">
          <vector name="origin" x="-13" y="23"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="5" width="29" height="7">
          <vector name="origin" x="-43" y="9"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="6" width="54" height="3">
          <vector name="origin" x="-41" y="8"/>
          <int name="delay" value="120"/>
        </canvas>
      </imgdir>
      <imgdir name="effect0">
        <canvas name="0" width="118" height="81">
          <vector name="origin" x="33" y="73"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="1" width="108" height="75">
          <vector name="origin" x="30" y="70"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="2" width="94" height="68">
          <vector name="origin" x="26" y="66"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="71" height="54">
          <vector name="origin" x="21" y="60"/>
          <int name="delay" value="120"/>
        </canvas>
        <int name="z" value="-1"/>
      </imgdir>
      <imgdir name="special">
        <canvas name="0" width="34" height="21">
          <vector name="origin" x="2" y="18"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="1" width="42" height="27">
          <vector name="origin" x="-3" y="24"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="2" width="40" height="26">
          <vector name="origin" x="1" y="23"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="40" height="24">
          <vector name="origin" x="-4" y="21"/>
          <int name="delay" value="120"/>
        </canvas>
      </imgdir>
      <int name="invisible" value="1"/>
      <int name="timeLimited" value="1"/>
      <int name="disable" value="1"/>
    </imgdir>
    <imgdir name="0001015">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="x" value="1"/>
          <int name="time" value="1"/>
          <string name="dateExpire" value="2009060800"/>
        </imgdir>
      </imgdir>
      <imgdir name="action">
        <string name="0" value="float"/>
      </imgdir>
      <int name="invisible" value="1"/>
      <int name="timeLimited" value="1"/>
      <int name="disable" value="1"/>
    </imgdir>
    <imgdir name="0001017">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="mpCon" value="10"/>
          <int name="time" value="2100000"/>
          <int name="pdd" value="10"/>
          <int name="mdd" value="10"/>
          <int name="x" value="1"/>
          <string name="dateExpire" value="2009090700"/>
        </imgdir>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="52" height="14">
          <vector name="origin" x="27" y="12"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="1" width="58" height="18">
          <vector name="origin" x="30" y="15"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="60" height="21">
          <vector name="origin" x="31" y="18"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="60" height="60">
          <vector name="origin" x="31" y="57"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="71" height="118">
          <vector name="origin" x="38" y="115"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="144" height="115">
          <vector name="origin" x="76" y="109"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="147" height="119">
          <vector name="origin" x="77" y="111"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="150" height="122">
          <vector name="origin" x="78" y="112"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="150" height="124">
          <vector name="origin" x="76" y="113"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="140" height="99">
          <vector name="origin" x="75" y="96"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="142" height="99">
          <vector name="origin" x="75" y="93"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="147" height="101">
          <vector name="origin" x="77" y="90"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="12" width="148" height="100">
          <vector name="origin" x="76" y="85"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="13" width="146" height="94">
          <vector name="origin" x="73" y="77"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="14" width="130" height="62">
          <vector name="origin" x="74" y="63"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="15" width="35" height="37">
          <vector name="origin" x="-22" y="40"/>
          <int name="delay" value="60"/>
        </canvas>
      </imgdir>
      <imgdir name="effect0">
        <canvas name="0" width="1" height="1">
          <vector name="origin" x="0" y="0"/>
          <int name="delay" value="300"/>
        </canvas>
        <canvas name="1" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="152" height="167">
          <vector name="origin" x="82" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="151" height="168">
          <vector name="origin" x="79" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="165" height="168">
          <vector name="origin" x="76" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="163" height="168">
          <vector name="origin" x="72" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="181" height="167">
          <vector name="origin" x="93" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="178" height="166">
          <vector name="origin" x="93" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="174" height="167">
          <vector name="origin" x="91" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <int name="z" value="-2"/>
      </imgdir>
      <int name="invisible" value="1"/>
      <int name="timeLimited" value="1"/>
      <int name="disable" value="1"/>
    </imgdir>
    <imgdir name="0001020">
      <string name="info" value="파라오의 분노"/>
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="effect">
        <canvas name="0" width="235" height="184">
          <vector name="origin" x="127" y="262"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="1" width="231" height="182">
          <vector name="origin" x="125" y="260"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="2" width="220" height="176">
          <vector name="origin" x="117" y="254"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="193" height="161">
          <vector name="origin" x="99" y="244"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="4" width="158" height="153">
          <vector name="origin" x="79" y="227"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="5" width="142" height="154">
          <vector name="origin" x="69" y="221"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="6" width="142" height="162">
          <vector name="origin" x="69" y="221"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="7" width="148" height="148">
          <vector name="origin" x="73" y="236"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="8" width="153" height="147">
          <vector name="origin" x="75" y="235"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="159" height="140">
          <vector name="origin" x="77" y="232"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="163" height="133">
          <vector name="origin" x="77" y="225"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="122" height="116">
          <vector name="origin" x="60" y="220"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="12" width="117" height="111">
          <vector name="origin" x="57" y="217"/>
          <int name="delay" value="60"/>
        </canvas>
        <int name="z" value="-2"/>
      </imgdir>
      <imgdir name="action">
        <string name="0" value="pyramid"/>
      </imgdir>
      <imgdir name="hit">
        <imgdir name="0">
          <canvas name="0" width="104" height="111">
            <vector name="origin" x="56" y="59"/>
            <int name="z" value="0"/>
          </canvas>
          <canvas name="1" width="89" height="84">
            <vector name="origin" x="42" y="45"/>
            <int name="z" value="0"/>
          </canvas>
          <canvas name="2" width="100" height="82">
            <vector name="origin" x="48" y="40"/>
            <int name="z" value="0"/>
          </canvas>
          <canvas name="3" width="117" height="95">
            <vector name="origin" x="56" y="46"/>
            <int name="z" value="0"/>
          </canvas>
          <canvas name="4" width="126" height="105">
            <vector name="origin" x="59" y="51"/>
            <int name="z" value="0"/>
          </canvas>
        </imgdir>
      </imgdir>
      <imgdir name="level">
        <imgdir name="1">
          <int name="damagepc" value="100"/>
          <int name="mobCount" value="15"/>
          <vector name="lt" x="-1000" y="-300"/>
          <vector name="rb" x="1000" y="15"/>
        </imgdir>
      </imgdir>
      <int name="invisible" value="1"/>
      <imgdir name="screen">
        <canvas name="0" width="1" height="1">
          <vector name="origin" x="0" y="0"/>
          <int name="z" value="0"/>
          <int name="delay" value="240"/>
        </canvas>
        <canvas name="1" width="240" height="160">
          <vector name="origin" x="116" y="189"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="2" width="256" height="212">
          <vector name="origin" x="124" y="241"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="260" height="231">
          <vector name="origin" x="126" y="266"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="4" width="262" height="231">
          <vector name="origin" x="127" y="274"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="5" width="264" height="233">
          <vector name="origin" x="128" y="278"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="6" width="265" height="234">
          <vector name="origin" x="129" y="279"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="7" width="591" height="497">
          <vector name="origin" x="155" y="328"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="8" width="874" height="496">
          <vector name="origin" x="441" y="325"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="9" width="898" height="498">
          <vector name="origin" x="434" y="321"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="10" width="912" height="493">
          <vector name="origin" x="440" y="316"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="11" width="921" height="497">
          <vector name="origin" x="447" y="318"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="12" width="936" height="499">
          <vector name="origin" x="459" y="318"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="13" width="948" height="504">
          <vector name="origin" x="468" y="316"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="14" width="985" height="502">
          <vector name="origin" x="483" y="313"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="15" width="1008" height="492">
          <vector name="origin" x="508" y="304"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="16" width="914" height="481">
          <vector name="origin" x="503" y="293"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="17" width="818" height="481">
          <vector name="origin" x="498" y="293"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="18" width="635" height="417">
          <vector name="origin" x="404" y="237"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="19" width="468" height="404">
          <vector name="origin" x="307" y="239"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="20" width="292" height="399">
          <vector name="origin" x="225" y="231"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="21" width="88" height="87">
          <vector name="origin" x="41" y="226"/>
          <int name="z" value="0"/>
          <int name="delay" value="120"/>
        </canvas>
      </imgdir>
    </imgdir>
    <imgdir name="0009000">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="invisible" value="1"/>
      <int name="mobCode" value="1210100"/>
      <int name="disable" value="1"/>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <string name="damage" value="150"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="0009001">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="invisible" value="1"/>
      <int name="mobCode" value="130100"/>
      <int name="disable" value="1"/>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <string name="damage" value="150"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="0009002">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="invisible" value="1"/>
      <int name="mobCode" value="210100"/>
      <int name="disable" value="1"/>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <string name="damage" value="150"/>
        </imgdir>
      </imgdir>
    </imgdir>
    <imgdir name="0001018">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="mpCon" value="10"/>
          <int name="time" value="2100000"/>
          <int name="pdd" value="10"/>
          <int name="mdd" value="10"/>
          <int name="x" value="1"/>
        </imgdir>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="52" height="14">
          <vector name="origin" x="27" y="12"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="1" width="58" height="18">
          <vector name="origin" x="30" y="15"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="60" height="21">
          <vector name="origin" x="31" y="18"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="60" height="60">
          <vector name="origin" x="31" y="57"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="71" height="118">
          <vector name="origin" x="38" y="115"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="144" height="115">
          <vector name="origin" x="76" y="109"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="147" height="119">
          <vector name="origin" x="77" y="111"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="150" height="122">
          <vector name="origin" x="78" y="112"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="150" height="124">
          <vector name="origin" x="76" y="113"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="140" height="99">
          <vector name="origin" x="75" y="96"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="142" height="99">
          <vector name="origin" x="75" y="93"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="147" height="101">
          <vector name="origin" x="77" y="90"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="12" width="148" height="100">
          <vector name="origin" x="76" y="85"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="13" width="146" height="94">
          <vector name="origin" x="73" y="77"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="14" width="130" height="62">
          <vector name="origin" x="74" y="63"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="15" width="35" height="37">
          <vector name="origin" x="-22" y="40"/>
          <int name="delay" value="60"/>
        </canvas>
      </imgdir>
      <imgdir name="effect0">
        <canvas name="0" width="1" height="1">
          <vector name="origin" x="0" y="0"/>
          <int name="delay" value="300"/>
        </canvas>
        <canvas name="1" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="152" height="167">
          <vector name="origin" x="82" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="151" height="168">
          <vector name="origin" x="79" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="165" height="168">
          <vector name="origin" x="76" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="163" height="168">
          <vector name="origin" x="72" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="181" height="167">
          <vector name="origin" x="93" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="178" height="166">
          <vector name="origin" x="93" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="174" height="167">
          <vector name="origin" x="91" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <int name="z" value="-2"/>
      </imgdir>
      <int name="invisible" value="1"/>
      <int name="timeLimited" value="1"/>
      <int name="disable" value="1"/>
    </imgdir>
    <imgdir name="0001019">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="mpCon" value="10"/>
          <int name="time" value="2100000"/>
          <int name="pdd" value="10"/>
          <int name="mdd" value="10"/>
          <int name="x" value="1"/>
        </imgdir>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="52" height="14">
          <vector name="origin" x="27" y="12"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="1" width="58" height="18">
          <vector name="origin" x="30" y="15"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="60" height="21">
          <vector name="origin" x="31" y="18"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="60" height="60">
          <vector name="origin" x="31" y="57"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="71" height="118">
          <vector name="origin" x="38" y="115"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="144" height="115">
          <vector name="origin" x="76" y="109"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="147" height="119">
          <vector name="origin" x="77" y="111"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="150" height="122">
          <vector name="origin" x="78" y="112"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="150" height="124">
          <vector name="origin" x="76" y="113"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="140" height="99">
          <vector name="origin" x="75" y="96"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="142" height="99">
          <vector name="origin" x="75" y="93"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="147" height="101">
          <vector name="origin" x="77" y="90"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="12" width="148" height="100">
          <vector name="origin" x="76" y="85"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="13" width="146" height="94">
          <vector name="origin" x="73" y="77"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="14" width="130" height="62">
          <vector name="origin" x="74" y="63"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="15" width="35" height="37">
          <vector name="origin" x="-22" y="40"/>
          <int name="delay" value="60"/>
        </canvas>
      </imgdir>
      <imgdir name="effect0">
        <canvas name="0" width="1" height="1">
          <vector name="origin" x="0" y="0"/>
          <int name="delay" value="300"/>
        </canvas>
        <canvas name="1" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="152" height="167">
          <vector name="origin" x="82" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="151" height="168">
          <vector name="origin" x="79" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="165" height="168">
          <vector name="origin" x="76" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="163" height="168">
          <vector name="origin" x="72" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="181" height="167">
          <vector name="origin" x="93" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="178" height="166">
          <vector name="origin" x="93" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="174" height="167">
          <vector name="origin" x="91" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <int name="z" value="-2"/>
      </imgdir>
      <int name="invisible" value="1"/>
      <int name="timeLimited" value="1"/>
      <int name="disable" value="1"/>
    </imgdir>
    <imgdir name="0001031">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="mpCon" value="10"/>
          <int name="time" value="2100000"/>
          <int name="pdd" value="10"/>
          <int name="mdd" value="10"/>
          <int name="x" value="1"/>
        </imgdir>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="52" height="14">
          <vector name="origin" x="27" y="12"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="1" width="58" height="18">
          <vector name="origin" x="30" y="15"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="60" height="21">
          <vector name="origin" x="31" y="18"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="60" height="60">
          <vector name="origin" x="31" y="57"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="71" height="118">
          <vector name="origin" x="38" y="115"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="144" height="115">
          <vector name="origin" x="76" y="109"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="147" height="119">
          <vector name="origin" x="77" y="111"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="150" height="122">
          <vector name="origin" x="78" y="112"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="150" height="124">
          <vector name="origin" x="76" y="113"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="140" height="99">
          <vector name="origin" x="75" y="96"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="142" height="99">
          <vector name="origin" x="75" y="93"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="147" height="101">
          <vector name="origin" x="77" y="90"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="12" width="148" height="100">
          <vector name="origin" x="76" y="85"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="13" width="146" height="94">
          <vector name="origin" x="73" y="77"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="14" width="130" height="62">
          <vector name="origin" x="74" y="63"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="15" width="35" height="37">
          <vector name="origin" x="-22" y="40"/>
          <int name="delay" value="60"/>
        </canvas>
      </imgdir>
      <imgdir name="effect0">
        <canvas name="0" width="1" height="1">
          <vector name="origin" x="0" y="0"/>
          <int name="delay" value="300"/>
        </canvas>
        <canvas name="1" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="2" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="3" width="152" height="167">
          <vector name="origin" x="82" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="4" width="151" height="168">
          <vector name="origin" x="79" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="5" width="165" height="168">
          <vector name="origin" x="76" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="6" width="163" height="168">
          <vector name="origin" x="72" y="168"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="7" width="181" height="167">
          <vector name="origin" x="93" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="8" width="178" height="166">
          <vector name="origin" x="93" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="9" width="174" height="167">
          <vector name="origin" x="91" y="167"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="10" width="167" height="166">
          <vector name="origin" x="89" y="166"/>
          <int name="delay" value="60"/>
        </canvas>
        <canvas name="11" width="161" height="165">
          <vector name="origin" x="86" y="165"/>
          <int name="delay" value="60"/>
        </canvas>
        <int name="z" value="-2"/>
      </imgdir>
      <int name="invisible" value="1"/>
      <int name="timeLimited" value="1"/>
      <int name="disable" value="1"/>
    </imgdir>
    <imgdir name="0000100">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconMouseOver" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="iconDisabled" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
        <int name="z" value="0"/>
      </canvas>
      <imgdir name="level">
        <imgdir name="1">
          <string name="hs" value="h1"/>
          <int name="time" value="2100000"/>
        </imgdir>
      </imgdir>
      <int name="disable" value="1"/>
      <int name="invisible" value="1"/>
    </imgdir>
  </imgdir>
</imgdir>
`

func Identity[M any](m M) M {
	return m
}

func TestReader(t *testing.T) {
	l, _ := test.NewNullLogger()

	rms := Read(l)(xml.FromByteArrayProvider([]byte(testXML)))
	rmm, err := model.CollectToMap[RestModel, uint32, RestModel](rms, RestModel.GetId, Identity)()
	if err != nil {
		t.Fatal(err)
	}
	if len(rmm) != 25 {
		t.Fatalf("len(rmm) = %d, want 25", len(rmm))
	}

	var rm RestModel
	var ef effect.RestModel
	var ok bool

	if rm, ok = rmm[1000]; !ok {
		t.Fatal("rmm[1000] does not exist.")
	}
	if len(rm.Effects) != 3 {
		t.Fatalf("len(rm.Effects) = %d, want 3", len(rm.Effects))
	}
	ef = rm.Effects[0]
	if ef.MPConsume != 3 {
		t.Fatalf("rm.Effects[0].MPConsume = %d, want 3", ef.MPConsume)
	}
	if ef.FixDamage != 10 {
		t.Fatalf("rm.Effects[0].FixDamage = %d, want 10", ef.FixDamage)
	}
	ef = rm.Effects[1]
	if ef.MPConsume != 5 {
		t.Fatalf("rm.Effects[1].MPConsume = %d, want 5", ef.MPConsume)
	}
	if ef.FixDamage != 25 {
		t.Fatalf("rm.Effects[1].FixDamage = %d, want 25", ef.FixDamage)
	}
	ef = rm.Effects[2]
	if ef.MPConsume != 7 {
		t.Fatalf("rm.Effects[2].MPConsume = %d, want 7", ef.MPConsume)
	}
	if ef.FixDamage != 40 {
		t.Fatalf("rm.Effects[2].FixDamage = %d, want 40", ef.FixDamage)
	}

	if rm, ok = rmm[1001]; !ok {
		t.Fatal("rmm[1001] does not exist.")
	}
	if len(rm.Effects) != 3 {
		t.Fatalf("len(rm.Effects) = %d, want 3", len(rm.Effects))
	}
	ef = rm.Effects[0]
	if ef.MPConsume != 5 {
		t.Fatalf("rm.Effects[0].MPConsume = %d, want 5", ef.MPConsume)
	}
	if ef.Duration != 30 {
		t.Fatalf("rm.Effects[0].Duration = %d, want 30", ef.Duration)
	}
	if ef.X != 4 {
		t.Fatalf("rm.Effects[0].X = %d, want 4", ef.X)
	}
	if ef.Cooldown != 120 {
		t.Fatalf("rm.Effects[0].Cooldown = %d, want 120", ef.Cooldown)
	}
	ef = rm.Effects[1]
	if ef.MPConsume != 10 {
		t.Fatalf("rm.Effects[1].MPConsume = %d, want 10", ef.MPConsume)
	}
	if ef.Duration != 30 {
		t.Fatalf("rm.Effects[1].Duration = %d, want 30", ef.Duration)
	}
	if ef.X != 8 {
		t.Fatalf("rm.Effects[1].X = %d, want 8", ef.X)
	}
	if ef.Cooldown != 120 {
		t.Fatalf("rm.Effects[1].Cooldown = %d, want 120", ef.Cooldown)
	}
	ef = rm.Effects[2]
	if ef.MPConsume != 15 {
		t.Fatalf("rm.Effects[2].MPConsume = %d, want 15", ef.MPConsume)
	}
	if ef.Duration != 30 {
		t.Fatalf("rm.Effects[2].Duration = %d, want 30", ef.Duration)
	}
	if ef.X != 12 {
		t.Fatalf("rm.Effects[2].X = %d, want 12", ef.X)
	}
	if ef.Cooldown != 120 {
		t.Fatalf("rm.Effects[2].Cooldown = %d, want 120", ef.Cooldown)
	}

	if rm, ok = rmm[1002]; !ok {
		t.Fatal("rmm[1002] does not exist.")
	}
	if len(rm.Effects) != 3 {
		t.Fatalf("len(rm.Effects) = %d, want 3", len(rm.Effects))
	}
	ef = rm.Effects[0]
	if ef.MPConsume != 4 {
		t.Fatalf("rm.Effects[0].MPConsume = %d, want 4", ef.MPConsume)
	}
	if ef.Duration != 4 {
		t.Fatalf("rm.Effects[0].Duration = %d, want 4", ef.Duration)
	}
	if ef.Speed != 10 {
		t.Fatalf("rm.Effects[0].Speed = %d, want 10", ef.Speed)
	}
	if ef.Cooldown != 60 {
		t.Fatalf("rm.Effects[0].Cooldown = %d, want 60", ef.Cooldown)
	}
	ef = rm.Effects[1]
	if ef.MPConsume != 7 {
		t.Fatalf("rm.Effects[1].MPConsume = %d, want 7", ef.MPConsume)
	}
	if ef.Duration != 8 {
		t.Fatalf("rm.Effects[1].Duration = %d, want 8", ef.Duration)
	}
	if ef.Speed != 15 {
		t.Fatalf("rm.Effects[1].Speed = %d, want 15", ef.Speed)
	}
	if ef.Cooldown != 60 {
		t.Fatalf("rm.Effects[1].Cooldown = %d, want 60", ef.Cooldown)
	}
	ef = rm.Effects[2]
	if ef.MPConsume != 10 {
		t.Fatalf("rm.Effects[2].MPConsume = %d, want 10", ef.MPConsume)
	}
	if ef.Duration != 12 {
		t.Fatalf("rm.Effects[2].Duration = %d, want 12", ef.Duration)
	}
	if ef.Speed != 20 {
		t.Fatalf("rm.Effects[2].Speed = %d, want 20", ef.Speed)
	}
	if ef.Cooldown != 60 {
		t.Fatalf("rm.Effects[2].Cooldown = %d, want 60", ef.Cooldown)
	}
}

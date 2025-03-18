package monster

import (
	"atlas-data/xml"
	"context"
	"github.com/Chronicle20/atlas-tenant"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"strconv"
	"testing"
)

const testXML = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<imgdir name="8510000.img">
  <imgdir name="info">
    <int name="bodyAttack" value="1"/>
    <int name="level" value="110"/>
    <int name="maxHP" value="30000000"/>
    <int name="maxMP" value="3000000"/>
    <int name="speed" value="0"/>
    <int name="PADamage" value="790"/>
    <int name="PDDamage" value="1150"/>
    <int name="MADamage" value="780"/>
    <int name="MDDamage" value="1270"/>
    <int name="acc" value="245"/>
    <int name="eva" value="14"/>
    <int name="exp" value="1300000"/>
    <int name="undead" value="0"/>
    <int name="pushed" value="10000"/>
    <int name="hpRecovery" value="10000"/>
    <int name="mpRecovery" value="50000"/>
    <float name="fs" value="10.0"/>
    <string name="elemAttr" value="F3"/>
    <int name="summonType" value="1"/>
    <int name="hpTagColor" value="1"/>
    <int name="hpTagBgcolor" value="5"/>
    <int name="noFlip" value="1"/>
    <int name="boss" value="1"/>
    <int name="firstAttack" value="1"/>
    <int name="publicReward" value="1"/>
    <int name="rareItemDropLevel" value="2"/>
    <imgdir name="skill">
      <imgdir name="0">
        <int name="skill" value="114"/>
        <int name="action" value="3"/>
        <int name="level" value="5"/>
        <int name="effectAfter" value="0"/>
      </imgdir>
      <imgdir name="1">
        <int name="skill" value="200"/>
        <int name="action" value="1"/>
        <int name="level" value="41"/>
        <int name="effectAfter" value="0"/>
      </imgdir>
      <imgdir name="2">
        <int name="skill" value="127"/>
        <int name="action" value="2"/>
        <int name="level" value="2"/>
        <int name="effectAfter" value="0"/>
      </imgdir>
      <imgdir name="3">
        <int name="skill" value="140"/>
        <int name="action" value="4"/>
        <int name="level" value="5"/>
        <int name="effectAfter" value="0"/>
      </imgdir>
      <imgdir name="4">
        <int name="skill" value="141"/>
        <int name="action" value="4"/>
        <int name="level" value="4"/>
        <int name="effectAfter" value="0"/>
      </imgdir>
      <imgdir name="5">
        <int name="skill" value="120"/>
        <int name="action" value="3"/>
        <int name="level" value="5"/>
        <int name="effectAfter" value="0"/>
      </imgdir>
      <imgdir name="6">
        <int name="skill" value="200"/>
        <int name="action" value="1"/>
        <int name="level" value="42"/>
        <int name="effectAfter" value="0"/>
      </imgdir>
    </imgdir>
    <imgdir name="default">
      <canvas name="0" width="200" height="214">
        <vector name="origin" x="100" y="107"/>
        <int name="z" value="0"/>
      </canvas>
    </imgdir>
  </imgdir>
  <imgdir name="stand">
    <canvas name="0" width="416" height="364">
      <vector name="origin" x="217" y="336"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="1" width="417" height="365">
      <vector name="origin" x="218" y="335"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="2" width="418" height="368">
      <vector name="origin" x="219" y="336"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="3" width="419" height="371">
      <vector name="origin" x="220" y="337"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="2"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="4" width="418" height="368">
      <vector name="origin" x="219" y="336"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="1"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="5" width="417" height="365">
      <vector name="origin" x="218" y="335"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="6" width="416" height="364">
      <vector name="origin" x="217" y="336"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="7" width="417" height="365">
      <vector name="origin" x="218" y="335"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="8" width="418" height="368">
      <vector name="origin" x="219" y="336"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="9" width="419" height="371">
      <vector name="origin" x="220" y="337"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="10" width="418" height="368">
      <vector name="origin" x="219" y="336"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="11" width="417" height="365">
      <vector name="origin" x="218" y="335"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <uol name="12" value="0"/>
    <uol name="13" value="1"/>
    <uol name="14" value="2"/>
    <uol name="15" value="3"/>
    <uol name="16" value="4"/>
    <uol name="17" value="5"/>
    <uol name="18" value="6"/>
    <uol name="19" value="7"/>
    <uol name="20" value="8"/>
    <uol name="21" value="9"/>
    <uol name="22" value="10"/>
    <uol name="23" value="11"/>
    <uol name="24" value="0"/>
    <uol name="25" value="1"/>
    <uol name="26" value="2"/>
    <uol name="27" value="3"/>
    <uol name="28" value="4"/>
    <uol name="29" value="5"/>
    <uol name="30" value="6"/>
    <uol name="31" value="7"/>
    <uol name="32" value="8"/>
    <uol name="33" value="9"/>
    <uol name="34" value="10"/>
    <uol name="35" value="11"/>
  </imgdir>
  <imgdir name="skill1">
    <canvas name="0" width="415" height="364">
      <vector name="origin" x="217" y="336"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="1" width="416" height="368">
      <vector name="origin" x="218" y="338"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="2" width="417" height="379">
      <vector name="origin" x="219" y="347"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="3" width="418" height="393">
      <vector name="origin" x="220" y="359"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="1"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="4" width="417" height="402">
      <vector name="origin" x="219" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="1"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="5" width="416" height="399">
      <vector name="origin" x="218" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="6" width="415" height="398">
      <vector name="origin" x="217" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="7" width="416" height="399">
      <vector name="origin" x="218" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="8" width="417" height="402">
      <vector name="origin" x="219" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="9" width="418" height="403">
      <vector name="origin" x="220" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="10" width="417" height="402">
      <vector name="origin" x="219" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="2"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="11" width="416" height="377">
      <vector name="origin" x="218" y="347"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
  </imgdir>
  <imgdir name="skill2">
    <canvas name="0" width="415" height="364">
      <vector name="origin" x="217" y="336"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="95"/>
    </canvas>
    <canvas name="1" width="416" height="368">
      <vector name="origin" x="218" y="338"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="95"/>
    </canvas>
    <canvas name="2" width="417" height="379">
      <vector name="origin" x="219" y="347"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="95"/>
    </canvas>
    <canvas name="3" width="418" height="393">
      <vector name="origin" x="220" y="359"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="1"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="95"/>
    </canvas>
    <canvas name="4" width="417" height="402">
      <vector name="origin" x="219" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="1"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="95"/>
    </canvas>
    <canvas name="5" width="416" height="399">
      <vector name="origin" x="218" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="95"/>
    </canvas>
    <canvas name="6" width="415" height="398">
      <vector name="origin" x="217" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="150"/>
    </canvas>
    <canvas name="7" width="446" height="399">
      <vector name="origin" x="248" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="8" width="440" height="402">
      <vector name="origin" x="242" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="2"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="9" width="441" height="403">
      <vector name="origin" x="243" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="10" width="445" height="402">
      <vector name="origin" x="247" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="2"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="11" width="445" height="388">
      <vector name="origin" x="247" y="358"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
  </imgdir>
  <imgdir name="skill3">
    <canvas name="0" width="415" height="364">
      <vector name="origin" x="217" y="336"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="1" width="416" height="368">
      <vector name="origin" x="218" y="338"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="2" width="417" height="379">
      <vector name="origin" x="219" y="347"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="3" width="418" height="393">
      <vector name="origin" x="220" y="359"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="1"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="4" width="417" height="402">
      <vector name="origin" x="219" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="1"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="5" width="432" height="399">
      <vector name="origin" x="234" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="6" width="426" height="398">
      <vector name="origin" x="228" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="7" width="418" height="399">
      <vector name="origin" x="220" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="8" width="422" height="402">
      <vector name="origin" x="224" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="9" width="418" height="403">
      <vector name="origin" x="220" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="10" width="417" height="402">
      <vector name="origin" x="219" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="2"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="11" width="416" height="377">
      <vector name="origin" x="218" y="347"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
  </imgdir>
  <imgdir name="skill4">
    <canvas name="0" width="415" height="364">
      <vector name="origin" x="217" y="336"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="1" width="416" height="368">
      <vector name="origin" x="218" y="338"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="2" width="418" height="379">
      <vector name="origin" x="220" y="347"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="3" width="420" height="393">
      <vector name="origin" x="222" y="359"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="1"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="4" width="421" height="402">
      <vector name="origin" x="223" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="1"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="5" width="421" height="399">
      <vector name="origin" x="223" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="6" width="421" height="398">
      <vector name="origin" x="223" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="7" width="420" height="399">
      <vector name="origin" x="222" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="8" width="419" height="402">
      <vector name="origin" x="221" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="2"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="9" width="418" height="403">
      <vector name="origin" x="220" y="369"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="10" width="417" height="402">
      <vector name="origin" x="219" y="370"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="196" y="2"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
    <canvas name="11" width="416" height="377">
      <vector name="origin" x="218" y="347"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="198" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="110"/>
    </canvas>
  </imgdir>
  <imgdir name="attack1">
    <imgdir name="info">
      <imgdir name="range">
        <vector name="lt" x="-950" y="-510"/>
        <vector name="rb" x="-60" y="65"/>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="243" height="315">
          <vector name="origin" x="260" y="269"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="1" width="249" height="377">
          <vector name="origin" x="266" y="334"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="2" width="251" height="409">
          <vector name="origin" x="266" y="370"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="257" height="398">
          <vector name="origin" x="271" y="360"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="4" width="275" height="404">
          <vector name="origin" x="288" y="370"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="5" width="288" height="401">
          <vector name="origin" x="301" y="367"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="6" width="296" height="411">
          <vector name="origin" x="308" y="372"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="7" width="303" height="419">
          <vector name="origin" x="315" y="378"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="8" width="250" height="416">
          <vector name="origin" x="262" y="371"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="9" width="247" height="417">
          <vector name="origin" x="261" y="375"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="10" width="238" height="376">
          <vector name="origin" x="257" y="339"/>
          <int name="delay" value="120"/>
        </canvas>
      </imgdir>
      <imgdir name="effect0">
        <int name="effectType" value="1"/>
        <int name="effectDistance" value="200"/>
        <int name="randomPos" value="1"/>
        <int name="delay" value="2000"/>
        <imgdir name="0">
          <canvas name="0" width="48" height="21">
            <vector name="origin" x="24" y="13"/>
            <int name="delay" value="100"/>
          </canvas>
          <canvas name="1" width="60" height="25">
            <vector name="origin" x="30" y="17"/>
            <int name="delay" value="100"/>
          </canvas>
          <canvas name="2" width="66" height="32">
            <vector name="origin" x="33" y="24"/>
            <int name="delay" value="100"/>
          </canvas>
          <canvas name="3" width="60" height="36">
            <vector name="origin" x="30" y="28"/>
            <int name="delay" value="300"/>
          </canvas>
          <canvas name="4" width="72" height="211">
            <vector name="origin" x="35" y="208"/>
            <int name="delay" value="100"/>
          </canvas>
          <canvas name="5" width="68" height="211">
            <vector name="origin" x="34" y="208"/>
            <int name="delay" value="100"/>
          </canvas>
          <canvas name="6" width="70" height="211">
            <vector name="origin" x="35" y="208"/>
            <int name="delay" value="100"/>
          </canvas>
          <uol name="7" value="4"/>
          <uol name="8" value="5"/>
          <uol name="9" value="6"/>
          <uol name="10" value="4"/>
          <uol name="11" value="5"/>
          <uol name="12" value="6"/>
          <uol name="13" value="4"/>
          <uol name="14" value="5"/>
          <uol name="15" value="6"/>
          <uol name="16" value="4"/>
          <uol name="17" value="5"/>
          <uol name="18" value="6"/>
          <uol name="19" value="4"/>
          <uol name="20" value="5"/>
          <uol name="21" value="6"/>
          <uol name="22" value="4"/>
          <uol name="23" value="5"/>
          <uol name="24" value="6"/>
          <uol name="25" value="4"/>
          <uol name="26" value="5"/>
          <uol name="27" value="6"/>
          <uol name="28" value="4"/>
          <uol name="29" value="5"/>
          <uol name="30" value="6"/>
          <uol name="31" value="4"/>
          <uol name="32" value="5"/>
          <uol name="33" value="6"/>
          <canvas name="34" width="51" height="154">
            <vector name="origin" x="24" y="175"/>
            <int name="delay" value="100"/>
          </canvas>
          <canvas name="35" width="54" height="154">
            <vector name="origin" x="28" y="175"/>
            <int name="delay" value="100"/>
          </canvas>
          <canvas name="36" width="59" height="145">
            <vector name="origin" x="27" y="180"/>
            <int name="delay" value="100"/>
          </canvas>
          <canvas name="37" width="48" height="120">
            <vector name="origin" x="27" y="173"/>
            <int name="delay" value="100"/>
          </canvas>
        </imgdir>
        <uol name="1" value="0"/>
        <uol name="2" value="0"/>
        <uol name="3" value="0"/>
        <uol name="4" value="0"/>
        <vector name="lt" x="-950" y="-510"/>
        <vector name="rb" x="-60" y="65"/>
      </imgdir>
      <imgdir name="hit">
        <canvas name="0" width="99" height="107">
          <vector name="origin" x="51" y="104"/>
          <int name="delay" value="70"/>
        </canvas>
        <canvas name="1" width="80" height="77">
          <vector name="origin" x="38" y="88"/>
          <int name="delay" value="70"/>
        </canvas>
      </imgdir>
      <int name="type" value="0"/>
      <int name="conMP" value="10"/>
      <int name="effectAfter" value="120"/>
      <int name="attackAfter" value="1000"/>
      <int name="magic" value="1"/>
      <string name="elemAttr" value="F"/>
      <int name="deadlyAttack" value="1"/>
      <int name="doFirst" value="1"/>
    </imgdir>
    <canvas name="0" width="416" height="366">
      <vector name="origin" x="217" y="338"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="1" width="416" height="370">
      <vector name="origin" x="217" y="340"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="2" width="416" height="379">
      <vector name="origin" x="217" y="347"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="3" width="416" height="386">
      <vector name="origin" x="217" y="352"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="4" width="416" height="388">
      <vector name="origin" x="217" y="354"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="5" width="416" height="386">
      <vector name="origin" x="217" y="354"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="6" width="402" height="386">
      <vector name="origin" x="203" y="354"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="7" width="402" height="386">
      <vector name="origin" x="203" y="354"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="8" width="402" height="388">
      <vector name="origin" x="203" y="354"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="9" width="402" height="388">
      <vector name="origin" x="203" y="354"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="10" width="402" height="384">
      <vector name="origin" x="203" y="352"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="11" width="416" height="370">
      <vector name="origin" x="217" y="341"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="197" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <uol name="12" value="../stand/0"/>
    <uol name="13" value="../stand/1"/>
    <uol name="14" value="../stand/2"/>
    <uol name="15" value="../stand/3"/>
    <uol name="16" value="../stand/4"/>
    <uol name="17" value="../stand/5"/>
    <uol name="18" value="../stand/6"/>
    <uol name="19" value="../stand/7"/>
    <uol name="20" value="../stand/8"/>
    <uol name="21" value="../stand/9"/>
    <uol name="22" value="../stand/10"/>
    <uol name="23" value="../stand/11"/>
    <uol name="24" value="../stand/0"/>
    <uol name="25" value="../stand/1"/>
    <uol name="26" value="../stand/2"/>
    <uol name="27" value="../stand/3"/>
    <uol name="28" value="../stand/4"/>
    <uol name="29" value="../stand/5"/>
    <uol name="30" value="../stand/6"/>
    <uol name="31" value="../stand/7"/>
    <uol name="32" value="../stand/8"/>
    <uol name="33" value="../stand/9"/>
    <uol name="34" value="../stand/10"/>
    <uol name="35" value="../stand/11"/>
  </imgdir>
  <imgdir name="attack2">
    <imgdir name="info">
      <imgdir name="range">
        <vector name="lt" x="-230" y="-150"/>
        <vector name="rb" x="-110" y="0"/>
        <int name="start" value="-6"/>
        <int name="areaCount" value="12"/>
        <int name="attackCount" value="7"/>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="289" height="292">
          <vector name="origin" x="315" y="276"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="1" width="290" height="357">
          <vector name="origin" x="310" y="336"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="2" width="288" height="390">
          <vector name="origin" x="304" y="368"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="280" height="382">
          <vector name="origin" x="294" y="359"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="4" width="277" height="392">
          <vector name="origin" x="290" y="368"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="5" width="278" height="389">
          <vector name="origin" x="291" y="365"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="6" width="279" height="395">
          <vector name="origin" x="291" y="371"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="7" width="280" height="401">
          <vector name="origin" x="292" y="377"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="8" width="281" height="394">
          <vector name="origin" x="293" y="371"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="9" width="277" height="392">
          <vector name="origin" x="291" y="373"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="10" width="272" height="353">
          <vector name="origin" x="290" y="336"/>
          <int name="delay" value="120"/>
        </canvas>
      </imgdir>
      <imgdir name="areaWarning">
        <canvas name="0" width="58" height="30">
          <vector name="origin" x="28" y="24"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="1" width="55" height="50">
          <vector name="origin" x="26" y="47"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="2" width="57" height="50">
          <vector name="origin" x="26" y="47"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="3" width="57" height="50">
          <vector name="origin" x="26" y="47"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="4" width="57" height="50">
          <vector name="origin" x="26" y="47"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="5" width="57" height="50">
          <vector name="origin" x="26" y="47"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="6" width="57" height="50">
          <vector name="origin" x="26" y="47"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="7" width="57" height="50">
          <vector name="origin" x="26" y="47"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="8" width="57" height="50">
          <vector name="origin" x="26" y="47"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="9" width="57" height="50">
          <vector name="origin" x="26" y="47"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="10" width="57" height="50">
          <vector name="origin" x="26" y="47"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="11" width="57" height="50">
          <vector name="origin" x="26" y="47"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="12" width="57" height="49">
          <vector name="origin" x="26" y="46"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="13" width="53" height="49">
          <vector name="origin" x="26" y="46"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="14" width="51" height="48">
          <vector name="origin" x="25" y="45"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="15" width="107" height="191">
          <vector name="origin" x="53" y="567"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="16" width="107" height="191">
          <vector name="origin" x="53" y="454"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="17" width="107" height="191">
          <vector name="origin" x="53" y="313"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="18" width="145" height="163">
          <vector name="origin" x="73" y="161"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="19" width="153" height="104">
          <vector name="origin" x="77" y="102"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="20" width="154" height="104">
          <vector name="origin" x="76" y="101"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="21" width="163" height="104">
          <vector name="origin" x="81" y="101"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="22" width="163" height="100">
          <vector name="origin" x="80" y="101"/>
          <int name="delay" value="130"/>
        </canvas>
        <canvas name="23" width="161" height="98">
          <vector name="origin" x="82" y="101"/>
          <int name="delay" value="130"/>
        </canvas>
        <canvas name="24" width="107" height="98">
          <vector name="origin" x="53" y="101"/>
          <int name="delay" value="130"/>
        </canvas>
        <canvas name="25" width="107" height="98">
          <vector name="origin" x="53" y="101"/>
          <int name="delay" value="130"/>
        </canvas>
      </imgdir>
      <imgdir name="hit">
        <canvas name="0" width="99" height="107">
          <vector name="origin" x="50" y="98"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="1" width="80" height="77">
          <vector name="origin" x="38" y="75"/>
          <int name="delay" value="100"/>
        </canvas>
        <int name="attach" value="1"/>
      </imgdir>
      <int name="type" value="3"/>
      <int name="conMP" value="10"/>
      <int name="effectAfter" value="120"/>
      <int name="attackAfter" value="1770"/>
      <int name="magic" value="1"/>
    </imgdir>
    <uol name="0" value="../attack1/0"/>
    <uol name="1" value="../attack1/1"/>
    <uol name="2" value="../attack1/2"/>
    <uol name="3" value="../attack1/3"/>
    <uol name="4" value="../attack1/4"/>
    <uol name="5" value="../attack1/5"/>
    <uol name="6" value="../attack1/6"/>
    <uol name="7" value="../attack1/7"/>
    <uol name="8" value="../attack1/8"/>
    <uol name="9" value="../attack1/9"/>
    <uol name="10" value="../attack1/10"/>
    <uol name="11" value="../attack1/11"/>
    <uol name="12" value="../stand/0"/>
    <uol name="13" value="../stand/1"/>
    <uol name="14" value="../stand/2"/>
    <uol name="15" value="../stand/3"/>
    <uol name="16" value="../stand/4"/>
    <uol name="17" value="../stand/5"/>
    <uol name="18" value="../stand/6"/>
    <uol name="19" value="../stand/7"/>
    <uol name="20" value="../stand/8"/>
    <uol name="21" value="../stand/9"/>
    <uol name="22" value="../stand/10"/>
    <uol name="23" value="../stand/11"/>
  </imgdir>
  <imgdir name="attack3">
    <imgdir name="info">
      <imgdir name="range">
        <vector name="lt" x="-815" y="-245"/>
        <vector name="rb" x="-10" y="60"/>
      </imgdir>
      <imgdir name="effect">
        <canvas name="0" width="333" height="402">
          <vector name="origin" x="353" y="339"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="1" width="350" height="437">
          <vector name="origin" x="366" y="371"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="2" width="357" height="443">
          <vector name="origin" x="371" y="362"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="3" width="368" height="461">
          <vector name="origin" x="381" y="371"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="4" width="374" height="461">
          <vector name="origin" x="387" y="368"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="5" width="861" height="406">
          <vector name="origin" x="873" y="374"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="6" width="861" height="456">
          <vector name="origin" x="873" y="380"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="7" width="861" height="450">
          <vector name="origin" x="873" y="374"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="8" width="861" height="456">
          <vector name="origin" x="873" y="380"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="9" width="861" height="450">
          <vector name="origin" x="873" y="374"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="10" width="858" height="451">
          <vector name="origin" x="872" y="376"/>
          <int name="delay" value="120"/>
        </canvas>
        <canvas name="11" width="852" height="413">
          <vector name="origin" x="870" y="338"/>
          <int name="delay" value="120"/>
        </canvas>
      </imgdir>
      <imgdir name="hit">
        <canvas name="0" width="69" height="68">
          <vector name="origin" x="31" y="84"/>
          <int name="delay" value="100"/>
        </canvas>
        <canvas name="1" width="85" height="77">
          <vector name="origin" x="38" y="93"/>
          <int name="delay" value="100"/>
        </canvas>
      </imgdir>
      <int name="type" value="0"/>
      <int name="conMP" value="10"/>
      <int name="effectAfter" value="0"/>
      <int name="attackAfter" value="600"/>
      <int name="magic" value="1"/>
    </imgdir>
    <uol name="0" value="../attack1/0"/>
    <uol name="1" value="../attack1/1"/>
    <uol name="2" value="../attack1/2"/>
    <uol name="3" value="../attack1/3"/>
    <uol name="4" value="../attack1/4"/>
    <uol name="5" value="../attack1/5"/>
    <uol name="6" value="../attack1/6"/>
    <uol name="7" value="../attack1/7"/>
    <uol name="8" value="../attack1/8"/>
    <uol name="9" value="../attack1/9"/>
    <uol name="10" value="../attack1/10"/>
    <uol name="11" value="../attack1/11"/>
    <uol name="12" value="../stand/0"/>
    <uol name="13" value="../stand/1"/>
    <uol name="14" value="../stand/2"/>
    <uol name="15" value="../stand/3"/>
    <uol name="16" value="../stand/4"/>
    <uol name="17" value="../stand/5"/>
    <uol name="18" value="../stand/6"/>
    <uol name="19" value="../stand/7"/>
    <uol name="20" value="../stand/8"/>
    <uol name="21" value="../stand/9"/>
    <uol name="22" value="../stand/10"/>
    <uol name="23" value="../stand/11"/>
    <uol name="35" value="../stand/11"/>
  </imgdir>
  <imgdir name="hit1">
    <canvas name="0" width="444" height="386">
      <vector name="origin" x="244" y="355"/>
      <vector name="lt" x="-192" y="-272"/>
      <vector name="rb" x="199" y="0"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="600"/>
    </canvas>
  </imgdir>
  <imgdir name="die1">
    <canvas name="0" width="444" height="396">
      <vector name="origin" x="244" y="355"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="1" width="466" height="396">
      <vector name="origin" x="266" y="354"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="2" width="467" height="396">
      <vector name="origin" x="267" y="353"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="3" width="467" height="396">
      <vector name="origin" x="267" y="352"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="4" width="467" height="396">
      <vector name="origin" x="267" y="351"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="5" width="379" height="298">
      <vector name="origin" x="193" y="250"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="6" width="379" height="298">
      <vector name="origin" x="193" y="246"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="7" width="379" height="298">
      <vector name="origin" x="193" y="239"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="120"/>
    </canvas>
    <canvas name="8" width="378" height="298">
      <vector name="origin" x="192" y="231"/>
      <vector name="head" x="-64" y="-269"/>
      <int name="delay" value="300"/>
      <int name="a0" value="255"/>
      <int name="a1" value="0"/>
    </canvas>
  </imgdir>
</imgdir>
`

func testTenant() tenant.Model {
	t, _ := tenant.Create(uuid.New(), "GMS", 83, 1)
	return t
}

func TestReader(t *testing.T) {
	tt := testTenant()
	l, _ := test.NewNullLogger()
	ctx := tenant.WithContext(context.Background(), tt)

	_, _ = GetMonsterStringRegistry().Add(tt, MonsterString{id: strconv.Itoa(8510000), name: "Pianus"})
	_, _ = GetMonsterGaugeRegistry().Add(tt, Gauge{id: strconv.Itoa(8510000), exists: true})

	rm, err := Read(l)(ctx)(xml.FromByteArrayProvider([]byte(testXML)))()
	if err != nil {
		t.Fatal(err)
	}

	if rm.Id != 8510000 {
		t.Errorf("Id mismatch: got %d, expected 8510000", rm.Id)
	}
	if rm.Name != "Pianus" {
		t.Errorf("Name mismatch: got %s, expected Pianus", rm.Name)
	}
	if rm.HP != 30000000 {
		t.Errorf("HP mismatch: got %d, expected 30000000", rm.HP)
	}
	if rm.MP != 3000000 {
		t.Errorf("MP mismatch: got %d, expected 3000000", rm.MP)
	}
	if rm.Experience != 1300000 {
		t.Errorf("Experience mismatch: got %d, expected 1300000", rm.Experience)
	}
	if rm.Level != 110 {
		t.Errorf("Level mismatch: got %d, expected 110", rm.Level)
	}
	if rm.WeaponAttack != 790 {
		t.Errorf("WeaponAttack mismatch: got %d, expected 790", rm.WeaponAttack)
	}
	if rm.WeaponDefense != 1150 {
		t.Errorf("WeaponDefense mismatch: got %d, expected 1150", rm.WeaponDefense)
	}
	if rm.MagicAttack != 780 {
		t.Errorf("MagicAttack mismatch: got %d, expected 780", rm.MagicAttack)
	}
	if rm.MagicDefense != 1270 {
		t.Errorf("MagicDefense mismatch: got %d, expected 1270", rm.MagicDefense)
	}
	if rm.Friendly != false {
		t.Errorf("Friendly mismatch: got %t, expected false", rm.Friendly)
	}
	if rm.RemoveAfter != 0 {
		t.Errorf("RemoveAfter mismatch: got %d, expected 0", rm.RemoveAfter)
	}
	if rm.Boss != true {
		t.Errorf("Boss mismatch: got %t, expected true", rm.Boss)
	}
	if rm.ExplosiveReward != false {
		t.Errorf("ExplosiveReward mismatch: got %t, expected false", rm.ExplosiveReward)
	}
	if rm.FFALoot != true {
		t.Errorf("FFALoot mismatch: got %t, expected true", rm.FFALoot)
	}
	if rm.Undead != false {
		t.Errorf("Undead mismatch: got %t, expected false", rm.Undead)
	}
	if rm.BuffToGive != 0 {
		t.Errorf("BuffToGive mismatch: got %d, expected 0", rm.BuffToGive)
	}
	if rm.CP != 0 {
		t.Errorf("CP mismatch: got %d, expected 0", rm.CP)
	}
	if rm.RemoveOnMiss != false {
		t.Errorf("RemoveOnMiss mismatch: got %t, expected false", rm.RemoveOnMiss)
	}
	if rm.Changeable != false {
		t.Errorf("Changeable mismatch: got %t, expected false", rm.Changeable)
	}
	if rm.FixedStance != 5 {
		t.Errorf("FixedStance mismatch: got %d, expected 5", rm.FixedStance)
	}
	if rm.FirstAttack != false {
		t.Errorf("FirstAttack mismatch: got %t, expected false", rm.FirstAttack)
	}
	if rm.DropPeriod != 0 {
		t.Errorf("DropPeriod mismatch: got %d, expected 0", rm.DropPeriod)
	}
	if rm.Banish != (banish{"", 0, ""}) {
		t.Errorf("Banish mismatch: got %+v, expected %+v", rm.Banish, banish{"", 0, ""})
	}
	if rm.SelfDestruction != (selfDestruction{0, 0, 0}) {
		t.Errorf("SelfDestruction mismatch: got %+v, expected %+v", rm.SelfDestruction, selfDestruction{0, 0, 0})
	}
	if rm.CoolDamage != (coolDamage{0, 0}) {
		t.Errorf("CoolDamage mismatch: got %+v, expected %+v", rm.CoolDamage, coolDamage{0, 0})
	}
	// Validate AnimationTimes map
	expectedAnimationTimes := map[string]uint32{
		"attack1": 1440, "attack2": 0, "attack3": 0, "die1": 1260, "hit1": 600,
		"skill1": 1320, "skill2": 1270, "skill3": 1320, "skill4": 1320, "stand": 1440,
	}
	for k, v := range expectedAnimationTimes {
		if rm.AnimationTimes[k] != v {
			t.Errorf("AnimationTimes[%s] mismatch: got %d, expected %d", k, rm.AnimationTimes[k], v)
		}
	}

	// Validate Resistances map
	if rm.Resistances["FIRE"] != "WEAK" {
		t.Errorf("Resistances[FIRE] mismatch: got %s, expected WEAK", rm.Resistances["FIRE"])
	}
	// Validate Skills slice
	expectedSkills := []skill{
		{114, 5}, {200, 41}, {127, 2}, {140, 5}, {141, 4}, {120, 5}, {200, 42},
	}
	if len(rm.Skills) != len(expectedSkills) {
		t.Errorf("Skills length mismatch: got %d, expected %d", len(rm.Skills), len(expectedSkills))
	} else {
		for i, skill := range expectedSkills {
			if rm.Skills[i] != skill {
				t.Errorf("Skills[%d] mismatch: got %+v, expected %+v", i, rm.Skills[i], skill)
			}
		}
	}

	// Validate Revives slice
	if len(rm.Revives) != 0 {
		t.Errorf("Revives mismatch: expected an empty slice, got %+v", rm.Revives)
	}

	// Validate Tag Colors
	if rm.TagColor != 1 {
		t.Errorf("TagColor mismatch: got %d, expected 1", rm.TagColor)
	}
	if rm.TagBackgroundColor != 5 {
		t.Errorf("TagBackgroundColor mismatch: got %d, expected 5", rm.TagBackgroundColor)
	}
}

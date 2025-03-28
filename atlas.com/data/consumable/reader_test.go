package consumable

import (
	"atlas-data/xml"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/sirupsen/logrus/hooks/test"
	"strconv"
	"testing"
)

const testXML = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<imgdir name="0200.img">
  <imgdir name="02000000">
    <imgdir name="info">
      <canvas name="icon" width="27" height="30">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <canvas name="iconRaw" width="27" height="27">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <int name="price" value="25"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="50"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000001">
    <imgdir name="info">
      <canvas name="icon" width="27" height="30">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <canvas name="iconRaw" width="27" height="27">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <int name="price" value="80"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="150"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000002">
    <imgdir name="info">
      <canvas name="icon" width="27" height="30">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <canvas name="iconRaw" width="27" height="27">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <int name="price" value="160"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="300"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000003">
    <imgdir name="info">
      <canvas name="icon" width="27" height="30">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <canvas name="iconRaw" width="27" height="27">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <int name="price" value="100"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mp" value="100"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000004">
    <imgdir name="info">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="1000"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hpR" value="50"/>
      <int name="mpR" value="50"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000005">
    <imgdir name="info">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="2500"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hpR" value="100"/>
      <int name="mpR" value="100"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000006">
    <imgdir name="info">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="310"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mp" value="300"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000007">
    <imgdir name="info">
      <canvas name="icon" width="33" height="28">
        <vector name="origin" x="0" y="28"/>
      </canvas>
      <canvas name="iconRaw" width="33" height="23">
        <vector name="origin" x="0" y="28"/>
      </canvas>
      <int name="price" value="25"/>
      <int name="slotMax" value="150"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="50"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000008">
    <imgdir name="info">
      <canvas name="icon" width="33" height="28">
        <vector name="origin" x="0" y="28"/>
      </canvas>
      <canvas name="iconRaw" width="33" height="23">
        <vector name="origin" x="0" y="28"/>
      </canvas>
      <int name="price" value="80"/>
      <int name="slotMax" value="150"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="150"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000009">
    <imgdir name="info">
      <canvas name="icon" width="33" height="28">
        <vector name="origin" x="0" y="28"/>
      </canvas>
      <canvas name="iconRaw" width="33" height="23">
        <vector name="origin" x="0" y="28"/>
      </canvas>
      <int name="price" value="160"/>
      <int name="slotMax" value="150"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="300"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000010">
    <imgdir name="info">
      <canvas name="icon" width="33" height="28">
        <vector name="origin" x="0" y="28"/>
      </canvas>
      <canvas name="iconRaw" width="33" height="23">
        <vector name="origin" x="0" y="28"/>
      </canvas>
      <int name="price" value="100"/>
      <int name="slotMax" value="150"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mp" value="100"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000011">
    <imgdir name="info">
      <canvas name="icon" width="31" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="31" height="29">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="310"/>
      <int name="slotMax" value="150"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mp" value="300"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000012">
    <imgdir name="info">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="1000"/>
      <int name="tradeBlock" value="1"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hpR" value="50"/>
      <int name="mpR" value="50"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000013">
    <imgdir name="info">
      <canvas name="icon" width="26" height="27">
        <vector name="origin" x="-4" y="27"/>
      </canvas>
      <canvas name="iconRaw" width="21" height="21">
        <vector name="origin" x="-5" y="27"/>
      </canvas>
      <int name="price" value="1"/>
      <int name="tradeBlock" value="1"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="40"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000014">
    <imgdir name="info">
      <canvas name="icon" width="26" height="27">
        <vector name="origin" x="-4" y="27"/>
      </canvas>
      <canvas name="iconRaw" width="21" height="21">
        <vector name="origin" x="-5" y="27"/>
      </canvas>
      <int name="price" value="1"/>
      <int name="tradeBlock" value="1"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mp" value="80"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000015">
    <imgdir name="info">
      <canvas name="icon" width="27" height="30">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <canvas name="iconRaw" width="27" height="27">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <int name="price" value="0"/>
      <int name="tradeBlock" value="1"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="150"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000016">
    <imgdir name="info">
      <canvas name="icon" width="27" height="30">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <canvas name="iconRaw" width="27" height="27">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <int name="price" value="0"/>
      <int name="tradeBlock" value="1"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="300"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000017">
    <imgdir name="info">
      <canvas name="icon" width="27" height="30">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <canvas name="iconRaw" width="27" height="27">
        <vector name="origin" x="-3" y="30"/>
      </canvas>
      <int name="tradeBlock" value="1"/>
      <int name="price" value="0"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mp" value="100"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000018">
    <imgdir name="info">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="tradeBlock" value="1"/>
      <int name="price" value="0"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mp" value="300"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000019">
    <imgdir name="info">
      <canvas name="icon" width="32" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="tradeBlock" value="1"/>
      <int name="price" value="0"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hpR" value="100"/>
      <int name="mpR" value="100"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000020">
    <imgdir name="info">
      <canvas name="icon" width="27" height="29">
        <vector name="origin" x="-3" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="27" height="27">
        <vector name="origin" x="-3" y="29"/>
      </canvas>
      <int name="tradeBlock" value="1"/>
      <int name="price" value="20"/>
      <int name="notSale" value="1"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="50"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000021">
    <imgdir name="info">
      <canvas name="icon" width="27" height="29">
        <vector name="origin" x="-3" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="27" height="27">
        <vector name="origin" x="-3" y="29"/>
      </canvas>
      <int name="price" value="80"/>
      <int name="tradeBlock" value="1"/>
      <int name="notSale" value="1"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mp" value="100"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000022">
    <imgdir name="info">
      <canvas name="icon" width="30" height="30">
        <vector name="origin" x="-1" y="30"/>
      </canvas>
      <canvas name="iconRaw" width="30" height="29">
        <vector name="origin" x="-1" y="30"/>
      </canvas>
      <int name="price" value="50"/>
      <int name="tradeBlock" value="1"/>
      <int name="notSale" value="1"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="100"/>
    </imgdir>
  </imgdir>
  <imgdir name="02000023">
    <imgdir name="info">
      <canvas name="icon" width="30" height="30">
        <vector name="origin" x="-1" y="30"/>
      </canvas>
      <canvas name="iconRaw" width="30" height="29">
        <vector name="origin" x="-1" y="30"/>
      </canvas>
      <int name="price" value="50"/>
      <int name="tradeBlock" value="1"/>
      <int name="notSale" value="1"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mp" value="50"/>
    </imgdir>
  </imgdir>
  <imgdir name="02001000">
    <imgdir name="info">
      <canvas name="icon" width="26" height="28">
        <vector name="origin" x="-4" y="28"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="24">
        <vector name="origin" x="-4" y="28"/>
      </canvas>
      <int name="price" value="1600"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="1000"/>
      <int name="mp" value="1000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02001001">
    <imgdir name="info">
      <canvas name="icon" width="31" height="32">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="31" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="1150"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="2000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02001002">
    <imgdir name="info">
      <canvas name="icon" width="27" height="31">
        <vector name="origin" x="-3" y="31"/>
      </canvas>
      <canvas name="iconRaw" width="26" height="30">
        <vector name="origin" x="-3" y="31"/>
      </canvas>
      <int name="price" value="2000"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mp" value="2000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002000">
    <imgdir name="info">
      <canvas name="icon" width="26" height="28">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="25">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <int name="price" value="250"/>
    </imgdir>
    <imgdir name="spec">
      <int name="eva" value="5"/>
      <int name="time" value="180000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002001">
    <imgdir name="info">
      <canvas name="icon" width="26" height="28">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="25">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <int name="price" value="250"/>
    </imgdir>
    <imgdir name="spec">
      <int name="speed" value="8"/>
      <int name="time" value="180000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002002">
    <imgdir name="info">
      <canvas name="icon" width="26" height="28">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="25">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <int name="price" value="250"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mad" value="5"/>
      <int name="time" value="180000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002003">
    <imgdir name="info">
      <canvas name="icon" width="26" height="28">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="25">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <int name="price" value="550"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mad" value="10"/>
      <int name="time" value="180000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002004">
    <imgdir name="info">
      <canvas name="icon" width="26" height="28">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="25">
        <vector name="origin" x="-4" y="29"/>
        <string name="desc" value="천연 꿀을 모아 담은 단지이다.#nHP를 약 50% 회복신킨다.#nMP를 약 50% 회복시킨다."/>
      </canvas>
      <int name="price" value="250"/>
    </imgdir>
    <imgdir name="spec">
      <int name="pad" value="5"/>
      <int name="time" value="180000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002005">
    <imgdir name="info">
      <canvas name="icon" width="26" height="28">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="25">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <int name="price" value="250"/>
    </imgdir>
    <imgdir name="spec">
      <int name="acc" value="5"/>
      <int name="time" value="300000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002006">
    <imgdir name="info">
      <canvas name="icon" width="28" height="28">
        <vector name="origin" x="-3" y="28"/>
      </canvas>
      <canvas name="iconRaw" width="28" height="26">
        <vector name="origin" x="-3" y="28"/>
      </canvas>
      <int name="price" value="300"/>
    </imgdir>
    <imgdir name="spec">
      <int name="pad" value="5"/>
      <int name="time" value="600000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002007">
    <imgdir name="info">
      <canvas name="icon" width="29" height="29">
        <vector name="origin" x="-3" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="29" height="27">
        <vector name="origin" x="-3" y="29"/>
      </canvas>
      <int name="price" value="300"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mad" value="5"/>
      <int name="time" value="600000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002008">
    <imgdir name="info">
      <canvas name="icon" width="28" height="26">
        <vector name="origin" x="-2" y="26"/>
      </canvas>
      <canvas name="iconRaw" width="28" height="23">
        <vector name="origin" x="-2" y="26"/>
      </canvas>
      <int name="price" value="300"/>
    </imgdir>
    <imgdir name="spec">
      <int name="acc" value="10"/>
      <int name="time" value="600000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002009">
    <imgdir name="info">
      <canvas name="icon" width="27" height="27">
        <vector name="origin" x="-4" y="27"/>
      </canvas>
      <canvas name="iconRaw" width="26" height="23">
        <vector name="origin" x="-4" y="27"/>
      </canvas>
      <int name="price" value="300"/>
    </imgdir>
    <imgdir name="spec">
      <int name="eva" value="10"/>
      <int name="time" value="600000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002010">
    <imgdir name="info">
      <canvas name="icon" width="27" height="29">
        <vector name="origin" x="-3" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="27" height="24">
        <vector name="origin" x="-3" y="29"/>
      </canvas>
      <int name="price" value="250"/>
    </imgdir>
    <imgdir name="spec">
      <int name="speed" value="10"/>
      <int name="time" value="600000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002011">
    <imgdir name="info">
      <canvas name="icon" width="26" height="29">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="24" height="27">
        <vector name="origin" x="-5" y="29"/>
      </canvas>
      <int name="price" value="600"/>
      <int name="slotMax" value="100"/>
    </imgdir>
    <imgdir name="spec">
      <int name="pdd" value="30"/>
      <int name="time" value="1800000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002015">
    <imgdir name="info">
      <canvas name="icon" width="28" height="30">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="28" height="28">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="1000"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hpR" value="90"/>
      <int name="mpR" value="90"/>
      <int name="pad" value="5"/>
      <int name="pdd" value="40"/>
      <int name="time" value="900000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002016">
    <imgdir name="info">
      <canvas name="icon" width="26" height="29">
        <vector name="origin" x="-4" y="28"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="26">
        <vector name="origin" x="-5" y="28"/>
      </canvas>
      <int name="price" value="1000"/>
    </imgdir>
    <imgdir name="spec">
      <int name="eva" value="15"/>
      <int name="acc" value="9"/>
      <int name="time" value="480000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002017">
    <imgdir name="info">
      <canvas name="icon" width="26" height="29">
        <vector name="origin" x="-4" y="28"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="26">
        <vector name="origin" x="-5" y="28"/>
      </canvas>
      <int name="price" value="1000"/>
    </imgdir>
    <imgdir name="spec">
      <int name="pad" value="12"/>
      <int name="time" value="480000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002018">
    <imgdir name="info">
      <canvas name="icon" width="26" height="29">
        <vector name="origin" x="-4" y="28"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="26">
        <vector name="origin" x="-5" y="28"/>
      </canvas>
      <int name="price" value="1000"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mad" value="20"/>
      <int name="time" value="480000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002019">
    <imgdir name="info">
      <canvas name="icon" width="26" height="29">
        <vector name="origin" x="-4" y="28"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="26">
        <vector name="origin" x="-5" y="28"/>
      </canvas>
      <int name="price" value="1000"/>
    </imgdir>
    <imgdir name="spec">
      <int name="eva" value="20"/>
      <int name="time" value="480000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002020">
    <imgdir name="info">
      <canvas name="icon" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="600"/>
      <int name="slotMax" value="150"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mpR" value="60"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002021">
    <imgdir name="info">
      <canvas name="icon" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="600"/>
      <int name="slotMax" value="150"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hpR" value="60"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002022">
    <imgdir name="info">
      <canvas name="icon" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="750"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hpR" value="40"/>
      <int name="mpR" value="40"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002023">
    <imgdir name="info">
      <canvas name="icon" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="1100"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hpR" value="75"/>
      <int name="mpR" value="75"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002024">
    <imgdir name="info">
      <canvas name="icon" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="500"/>
      <int name="slotMax" value="150"/>
    </imgdir>
    <imgdir name="spec">
      <int name="mp" value="1500"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002025">
    <imgdir name="info">
      <canvas name="icon" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="500"/>
      <int name="slotMax" value="150"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hp" value="1500"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002026">
    <imgdir name="info">
      <canvas name="icon" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <canvas name="iconRaw" width="32" height="31">
        <vector name="origin" x="0" y="32"/>
      </canvas>
      <int name="price" value="1100"/>
    </imgdir>
    <imgdir name="spec">
      <int name="hpR" value="75"/>
      <int name="mpR" value="75"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002027">
    <imgdir name="info">
      <canvas name="icon" width="36" height="31">
        <vector name="origin" x="2" y="30"/>
      </canvas>
      <canvas name="iconRaw" width="36" height="28">
        <vector name="origin" x="2" y="30"/>
      </canvas>
      <int name="price" value="80"/>
    </imgdir>
    <imgdir name="spec">
      <int name="eva" value="5"/>
      <int name="time" value="1200000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002028">
    <imgdir name="info">
      <canvas name="icon" width="33" height="33">
        <vector name="origin" x="1" y="31"/>
      </canvas>
      <canvas name="iconRaw" width="33" height="33">
        <vector name="origin" x="1" y="31"/>
      </canvas>
      <int name="price" value="160"/>
    </imgdir>
    <imgdir name="spec">
      <int name="pad" value="25"/>
      <int name="mad" value="60"/>
      <int name="time" value="1200000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002029">
    <imgdir name="info">
      <canvas name="icon" width="28" height="30">
        <vector name="origin" x="-2" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="28" height="29">
        <vector name="origin" x="-2" y="29"/>
      </canvas>
      <int name="price" value="25"/>
    </imgdir>
    <imgdir name="spec">
      <int name="pdd" value="100"/>
      <int name="mdd" value="100"/>
      <int name="time" value="600000"/>
    </imgdir>
  </imgdir>
  <imgdir name="02002030">
    <imgdir name="info">
      <canvas name="icon" width="26" height="28">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <canvas name="iconRaw" width="25" height="25">
        <vector name="origin" x="-4" y="29"/>
      </canvas>
      <int name="price" value="1"/>
      <int name="tradeBlock" value="1"/>
      <int name="notSale" value="1"/>
    </imgdir>
    <imgdir name="spec">
      <int name="speed" value="5"/>
      <int name="time" value="600000"/>
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
	rmm, err := model.CollectToMap[RestModel, string, RestModel](rms, RestModel.GetID, Identity)()
	if err != nil {
		t.Fatal(err)
	}
	if len(rmm) != 55 {
		t.Fatalf("len(rmm) = %d, want 55", len(rmm))
	}

	var rm RestModel
	var ok bool
	var spec int32

	if rm, ok = rmm[strconv.Itoa(2000000)]; !ok {
		t.Fatalf("rmm[2000000] does not exist.")
	}
	if rm.Price != 25 {
		t.Fatalf("rm.Price = %d, want 25", rm.Price)
	}
	if spec, ok = rm.Spec[SpecTypeHP]; !ok {
		t.Fatalf("rmm.Spec[SpecTypeHP] does not exist.")
	}
	if spec != 50 {
		t.Fatalf("rmm.Spec[SpecTypeHP] = %d, want 50", spec)
	}

	if rm, ok = rmm[strconv.Itoa(2000004)]; !ok {
		t.Fatalf("rmm[2000004] does not exist.")
	}
	if rm.Price != 1000 {
		t.Fatalf("rm.Price = %d, want 1000", rm.Price)
	}
	if spec, ok = rm.Spec[SpecTypeHPRecovery]; !ok {
		t.Fatalf("rmm.Spec[SpecTypeHPRecovery] does not exist.")
	}
	if spec != 50 {
		t.Fatalf("rmm.Spec[SpecTypeHPRecovery] = %d, want 50", spec)
	}
	if spec, ok = rm.Spec[SpecTypeMPRecovery]; !ok {
		t.Fatalf("rmm.Spec[SpecTypeMPRecovery] does not exist.")
	}
	if spec != 50 {
		t.Fatalf("rmm.Spec[SpecTypeMPRecovery] = %d, want 50", spec)
	}

	if rm, ok = rmm[strconv.Itoa(2000020)]; !ok {
		t.Fatalf("rmm[2000020] does not exist.")
	}
	if rm.Price != 20 {
		t.Fatalf("rm.Price = %d, want 20", rm.Price)
	}
	if !rm.TradeBlock {
		t.Fatalf("rm.TradeBlock = false")
	}
	if !rm.NotSale {
		t.Fatalf("rm.NotSale = false")
	}
	if spec, ok = rm.Spec[SpecTypeHP]; !ok {
		t.Fatalf("rmm.Spec[SpecTypeHP] does not exist.")
	}
	if spec != 50 {
		t.Fatalf("rmm.Spec[SpecTypeHP] = %d, want 50", spec)
	}

	if rm, ok = rmm[strconv.Itoa(2001000)]; !ok {
		t.Fatalf("rmm[2001000] does not exist.")
	}
	if rm.Price != 1600 {
		t.Fatalf("rm.Price = %d, want 1600", rm.Price)
	}
	if spec, ok = rm.Spec[SpecTypeHP]; !ok {
		t.Fatalf("rmm.Spec[SpecTypeHP] does not exist.")
	}
	if spec != 1000 {
		t.Fatalf("rmm.Spec[SpecTypeHP] = %d, want 1000", spec)
	}
	if spec, ok = rm.Spec[SpecTypeMP]; !ok {
		t.Fatalf("rmm.Spec[SpecTypeMP] does not exist.")
	}
	if spec != 1000 {
		t.Fatalf("rmm.Spec[SpecTypeMP] = %d, want 1000", spec)
	}

	if rm, ok = rmm[strconv.Itoa(2002000)]; !ok {
		t.Fatalf("rmm[02002000] does not exist.")
	}
	if rm.Price != 250 {
		t.Fatalf("rm.Price = %d, want 250", rm.Price)
	}
	if spec, ok = rm.Spec[SpecTypeEvasion]; !ok {
		t.Fatalf("rmm.Spec[SpecTypeEvasion] does not exist.")
	}
	if spec != 5 {
		t.Fatalf("rmm.Spec[SpecTypeEvasion] = %d, want 5", spec)
	}
	if spec, ok = rm.Spec[SpecTypeTime]; !ok {
		t.Fatalf("rmm.Spec[SpecTypeTime] does not exist.")
	}
	if spec != 180000 {
		t.Fatalf("rmm.Spec[SpecTypeTime] = %d, want 180000", spec)
	}
}

package reactor

import (
	"atlas-data/xml"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/sirupsen/logrus/hooks/test"
	"testing"
)

const testXML = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<imgdir name="1002000.img">
  <imgdir name="info">
    <string name="name" value="거대병아리"/>
  </imgdir>
  <imgdir name="0">
    <canvas name="0" width="158" height="211">
      <vector name="origin" x="79" y="105"/>
      <int name="z" value="0"/>
    </canvas>
    <imgdir name="hit">
      <canvas name="0" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="1" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="2" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="3" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="4" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
    </imgdir>
  </imgdir>
  <imgdir name="1">
    <canvas name="0" width="158" height="211">
      <vector name="origin" x="79" y="105"/>
      <int name="z" value="0"/>
    </canvas>
    <imgdir name="hit">
      <canvas name="0" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="1" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="2" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="3" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="4" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
    </imgdir>
  </imgdir>
  <imgdir name="2">
    <canvas name="0" width="158" height="211">
      <vector name="origin" x="79" y="105"/>
      <int name="z" value="0"/>
    </canvas>
    <imgdir name="hit">
      <canvas name="0" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="1" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="2" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="3" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="4" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
    </imgdir>
  </imgdir>
  <imgdir name="3">
    <canvas name="0" width="158" height="211">
      <vector name="origin" x="79" y="105"/>
      <int name="z" value="0"/>
    </canvas>
    <imgdir name="hit">
      <canvas name="0" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
    </imgdir>
  </imgdir>
  <imgdir name="4">
    <canvas name="0" width="158" height="211">
      <vector name="origin" x="79" y="105"/>
      <int name="z" value="0"/>
    </canvas>
    <imgdir name="hit">
      <canvas name="0" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="1" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="2" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="3" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="4" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="5" width="158" height="202">
        <vector name="origin" x="79" y="96"/>
        <int name="z" value="0"/>
        <int name="delay" value="1200"/>
      </canvas>
      <canvas name="6" width="158" height="201">
        <vector name="origin" x="79" y="95"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="7" width="158" height="202">
        <vector name="origin" x="79" y="96"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="8" width="158" height="203">
        <vector name="origin" x="79" y="97"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="9" width="158" height="202">
        <vector name="origin" x="79" y="96"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="10" width="158" height="202">
        <vector name="origin" x="79" y="96"/>
        <int name="z" value="0"/>
        <int name="delay" value="1200"/>
      </canvas>
      <canvas name="11" width="158" height="202">
        <vector name="origin" x="79" y="96"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="12" width="158" height="203">
        <vector name="origin" x="79" y="97"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="13" width="158" height="202">
        <vector name="origin" x="79" y="96"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="14" width="158" height="202">
        <vector name="origin" x="79" y="96"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="15" width="158" height="202">
        <vector name="origin" x="79" y="96"/>
        <int name="z" value="0"/>
        <int name="delay" value="1200"/>
      </canvas>
      <canvas name="16" width="158" height="201">
        <vector name="origin" x="79" y="95"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="17" width="158" height="212">
        <vector name="origin" x="79" y="106"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="18" width="158" height="222">
        <vector name="origin" x="79" y="116"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="19" width="158" height="227">
        <vector name="origin" x="79" y="121"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="20" width="158" height="237">
        <vector name="origin" x="79" y="131"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="21" width="158" height="242">
        <vector name="origin" x="79" y="136"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="22" width="158" height="247">
        <vector name="origin" x="79" y="141"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="23" width="158" height="262">
        <vector name="origin" x="79" y="156"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="24" width="158" height="277">
        <vector name="origin" x="79" y="171"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="25" width="158" height="297">
        <vector name="origin" x="79" y="191"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="26" width="158" height="317">
        <vector name="origin" x="79" y="211"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="27" width="158" height="337">
        <vector name="origin" x="79" y="231"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="28" width="158" height="357">
        <vector name="origin" x="79" y="251"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="29" width="158" height="377">
        <vector name="origin" x="79" y="271"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="30" width="158" height="397">
        <vector name="origin" x="79" y="291"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="31" width="158" height="417">
        <vector name="origin" x="79" y="311"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="32" width="158" height="437">
        <vector name="origin" x="79" y="331"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="33" width="158" height="457">
        <vector name="origin" x="79" y="351"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="34" width="158" height="477">
        <vector name="origin" x="79" y="371"/>
        <int name="z" value="0"/>
      </canvas>
    </imgdir>
  </imgdir>
  <imgdir name="5">
    <canvas name="0" width="158" height="133">
      <vector name="origin" x="79" y="27"/>
      <int name="z" value="0"/>
    </canvas>
    <imgdir name="event">
      <imgdir name="0">
        <int name="type" value="101"/>
        <int name="state" value="0"/>
      </imgdir>
      <int name="timeout" value="1000"/>
    </imgdir>
    <int name="timeout" value="1000"/>
    <imgdir name="hit">
      <canvas name="0" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="1" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="2" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="3" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
      <canvas name="4" width="158" height="211">
        <vector name="origin" x="79" y="105"/>
        <int name="z" value="0"/>
      </canvas>
    </imgdir>
  </imgdir>
  <string name="action" value="babyBirdItem0"/>
</imgdir>
`

const linkTestXML = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<imgdir name="1020008.img">
  <imgdir name="info">
    <string name="info" value="91020002"/>
    <string name="link" value="1020000"/>
  </imgdir>
  <string name="action" value="s4hitmanMap8"/>
</imgdir>
`

var fixedNodeProvider = func(path string, id uint32) model.Provider[xml.Node] {
	return xml.FromByteArrayProvider([]byte(testXML))
}

var linkedNodeProvider = func(path string, id uint32) model.Provider[xml.Node] {
	if id == 0 {
		return xml.FromByteArrayProvider([]byte(linkTestXML))
	} else {
		return xml.FromByteArrayProvider([]byte(testXML))
	}
}

func TestReader(t *testing.T) {
	l, _ := test.NewNullLogger()

	rm, err := Read(l)("", 0, fixedNodeProvider)()
	if err != nil {
		t.Fatal(err)
	}
	if rm.Id != 1002000 {
		t.Fatal("id != 1002000")
	}
	if len(rm.StateInfo) != 6 {
		t.Fatal("len(rm.StateInfo) != 6")
	}
}

func TestLinkedReader(t *testing.T) {
	l, _ := test.NewNullLogger()

	rm, err := Read(l)("", 0, linkedNodeProvider)()
	if err != nil {
		t.Fatal(err)
	}
	if rm.Id != 1020008 {
		t.Fatal("id != 1020008")
	}
	if len(rm.StateInfo) != 6 {
		t.Fatal("len(rm.StateInfo) != 6")
	}
}

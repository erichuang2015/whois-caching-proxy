package main

//
// From: https://publicsuffix.org/list/public_suffix_list.dat
//		** REMOVED the private domains from the list
var publicsuffixes = `// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Please pull this list from, and only from https://publicsuffix.org/list/public_suffix_list.dat,
// rather than any other VCS sites. Pulling from any other URL is not guaranteed to be supported.

// Instructions on pulling and using this list can be found at https://publicsuffix.org/list/.

// ===BEGIN ICANN DOMAINS===

// ac : https://en.wikipedia.org/wiki/.ac
ac
com.ac
edu.ac
gov.ac
net.ac
mil.ac
org.ac

// ad : https://en.wikipedia.org/wiki/.ad
ad
nom.ad

// ae : https://en.wikipedia.org/wiki/.ae
// see also: "Domain Name Eligibility Policy" at http://www.aeda.ae/eng/aepolicy.php
ae
co.ae
net.ae
org.ae
sch.ae
ac.ae
gov.ae
mil.ae

// aero : see https://www.information.aero/index.php?id=66
aero
accident-investigation.aero
accident-prevention.aero
aerobatic.aero
aeroclub.aero
aerodrome.aero
agents.aero
aircraft.aero
airline.aero
airport.aero
air-surveillance.aero
airtraffic.aero
air-traffic-control.aero
ambulance.aero
amusement.aero
association.aero
author.aero
ballooning.aero
broker.aero
caa.aero
cargo.aero
catering.aero
certification.aero
championship.aero
charter.aero
civilaviation.aero
club.aero
conference.aero
consultant.aero
consulting.aero
control.aero
council.aero
crew.aero
design.aero
dgca.aero
educator.aero
emergency.aero
engine.aero
engineer.aero
entertainment.aero
equipment.aero
exchange.aero
express.aero
federation.aero
flight.aero
freight.aero
fuel.aero
gliding.aero
government.aero
groundhandling.aero
group.aero
hanggliding.aero
homebuilt.aero
insurance.aero
journal.aero
journalist.aero
leasing.aero
logistics.aero
magazine.aero
maintenance.aero
media.aero
microlight.aero
modelling.aero
navigation.aero
parachuting.aero
paragliding.aero
passenger-association.aero
pilot.aero
press.aero
production.aero
recreation.aero
repbody.aero
res.aero
research.aero
rotorcraft.aero
safety.aero
scientist.aero
services.aero
show.aero
skydiving.aero
software.aero
student.aero
trader.aero
trading.aero
trainer.aero
union.aero
workinggroup.aero
works.aero

// af : http://www.nic.af/help.jsp
af
gov.af
com.af
org.af
net.af
edu.af

// ag : http://www.nic.ag/prices.htm
ag
com.ag
org.ag
net.ag
co.ag
nom.ag

// ai : http://nic.com.ai/
ai
off.ai
com.ai
net.ai
org.ai

// al : http://www.ert.gov.al/ert_alb/faq_det.html?Id=31
al
com.al
edu.al
gov.al
mil.al
net.al
org.al

// am : https://en.wikipedia.org/wiki/.am
am

// ao : https://en.wikipedia.org/wiki/.ao
// http://www.dns.ao/REGISTR.DOC
ao
ed.ao
gv.ao
og.ao
co.ao
pb.ao
it.ao

// aq : https://en.wikipedia.org/wiki/.aq
aq

// ar : https://nic.ar/nic-argentina/normativa-vigente
ar
com.ar
edu.ar
gob.ar
gov.ar
int.ar
mil.ar
musica.ar
net.ar
org.ar
tur.ar

// arpa : https://en.wikipedia.org/wiki/.arpa
// Confirmed by registry <iana-questions@icann.org> 2008-06-18
arpa
e164.arpa
in-addr.arpa
ip6.arpa
iris.arpa
uri.arpa
urn.arpa

// as : https://en.wikipedia.org/wiki/.as
as
gov.as

// asia : https://en.wikipedia.org/wiki/.asia
asia

// at : https://en.wikipedia.org/wiki/.at
// Confirmed by registry <it@nic.at> 2008-06-17
at
ac.at
co.at
gv.at
or.at

// au : https://en.wikipedia.org/wiki/.au
// http://www.auda.org.au/
au
// 2LDs
com.au
net.au
org.au
edu.au
gov.au
asn.au
id.au
// Historic 2LDs (closed to new registration, but sites still exist)
info.au
conf.au
oz.au
// CGDNs - http://www.cgdn.org.au/
act.au
nsw.au
nt.au
qld.au
sa.au
tas.au
vic.au
wa.au
// 3LDs
act.edu.au
nsw.edu.au
nt.edu.au
qld.edu.au
sa.edu.au
tas.edu.au
vic.edu.au
wa.edu.au
// act.gov.au  Bug 984824 - Removed at request of Greg Tankard
// nsw.gov.au  Bug 547985 - Removed at request of <Shae.Donelan@services.nsw.gov.au>
// nt.gov.au  Bug 940478 - Removed at request of Greg Connors <Greg.Connors@nt.gov.au>
qld.gov.au
sa.gov.au
tas.gov.au
vic.gov.au
wa.gov.au

// aw : https://en.wikipedia.org/wiki/.aw
aw
com.aw

// ax : https://en.wikipedia.org/wiki/.ax
ax

// az : https://en.wikipedia.org/wiki/.az
az
com.az
net.az
int.az
gov.az
org.az
edu.az
info.az
pp.az
mil.az
name.az
pro.az
biz.az

// ba : http://nic.ba/users_data/files/pravilnik_o_registraciji.pdf
ba
com.ba
edu.ba
gov.ba
mil.ba
net.ba
org.ba

// bb : https://en.wikipedia.org/wiki/.bb
bb
biz.bb
co.bb
com.bb
edu.bb
gov.bb
info.bb
net.bb
org.bb
store.bb
tv.bb

// bd : https://en.wikipedia.org/wiki/.bd
*.bd

// be : https://en.wikipedia.org/wiki/.be
// Confirmed by registry <tech@dns.be> 2008-06-08
be
ac.be

// bf : https://en.wikipedia.org/wiki/.bf
bf
gov.bf

// bg : https://en.wikipedia.org/wiki/.bg
// https://www.register.bg/user/static/rules/en/index.html
bg
a.bg
b.bg
c.bg
d.bg
e.bg
f.bg
g.bg
h.bg
i.bg
j.bg
k.bg
l.bg
m.bg
n.bg
o.bg
p.bg
q.bg
r.bg
s.bg
t.bg
u.bg
v.bg
w.bg
x.bg
y.bg
z.bg
0.bg
1.bg
2.bg
3.bg
4.bg
5.bg
6.bg
7.bg
8.bg
9.bg

// bh : https://en.wikipedia.org/wiki/.bh
bh
com.bh
edu.bh
net.bh
org.bh
gov.bh

// bi : https://en.wikipedia.org/wiki/.bi
// http://whois.nic.bi/
bi
co.bi
com.bi
edu.bi
or.bi
org.bi

// biz : https://en.wikipedia.org/wiki/.biz
biz

// bj : https://en.wikipedia.org/wiki/.bj
bj
asso.bj
barreau.bj
gouv.bj

// bm : http://www.bermudanic.bm/dnr-text.txt
bm
com.bm
edu.bm
gov.bm
net.bm
org.bm

// bn : http://www.bnnic.bn/faqs
bn
com.bn
edu.bn
gov.bn
net.bn
org.bn

// bo : https://nic.bo/delegacion2015.php#h-1.10
bo
com.bo
edu.bo
gob.bo
int.bo
org.bo
net.bo
mil.bo
tv.bo
web.bo
// Social Domains
academia.bo
agro.bo
arte.bo
blog.bo
bolivia.bo
ciencia.bo
cooperativa.bo
democracia.bo
deporte.bo
ecologia.bo
economia.bo
empresa.bo
indigena.bo
industria.bo
info.bo
medicina.bo
movimiento.bo
musica.bo
natural.bo
nombre.bo
noticias.bo
patria.bo
politica.bo
profesional.bo
plurinacional.bo
pueblo.bo
revista.bo
salud.bo
tecnologia.bo
tksat.bo
transporte.bo
wiki.bo

// br : http://registro.br/dominio/categoria.html
// Submitted by registry <fneves@registro.br>
br
9guacu.br
abc.br
adm.br
adv.br
agr.br
aju.br
am.br
anani.br
aparecida.br
arq.br
art.br
ato.br
b.br
barueri.br
belem.br
bhz.br
bio.br
blog.br
bmd.br
boavista.br
bsb.br
campinagrande.br
campinas.br
caxias.br
cim.br
cng.br
cnt.br
com.br
contagem.br
coop.br
cri.br
cuiaba.br
curitiba.br
def.br
ecn.br
eco.br
edu.br
emp.br
eng.br
esp.br
etc.br
eti.br
far.br
feira.br
flog.br
floripa.br
fm.br
fnd.br
fortal.br
fot.br
foz.br
fst.br
g12.br
ggf.br
goiania.br
gov.br
// gov.br 26 states + df https://en.wikipedia.org/wiki/States_of_Brazil
ac.gov.br
al.gov.br
am.gov.br
ap.gov.br
ba.gov.br
ce.gov.br
df.gov.br
es.gov.br
go.gov.br
ma.gov.br
mg.gov.br
ms.gov.br
mt.gov.br
pa.gov.br
pb.gov.br
pe.gov.br
pi.gov.br
pr.gov.br
rj.gov.br
rn.gov.br
ro.gov.br
rr.gov.br
rs.gov.br
sc.gov.br
se.gov.br
sp.gov.br
to.gov.br
gru.br
imb.br
ind.br
inf.br
jab.br
jampa.br
jdf.br
joinville.br
jor.br
jus.br
leg.br
lel.br
londrina.br
macapa.br
maceio.br
manaus.br
maringa.br
mat.br
med.br
mil.br
morena.br
mp.br
mus.br
natal.br
net.br
niteroi.br
*.nom.br
not.br
ntr.br
odo.br
ong.br
org.br
osasco.br
palmas.br
poa.br
ppg.br
pro.br
psc.br
psi.br
pvh.br
qsl.br
radio.br
rec.br
recife.br
ribeirao.br
rio.br
riobranco.br
riopreto.br
salvador.br
sampa.br
santamaria.br
santoandre.br
saobernardo.br
saogonca.br
sjc.br
slg.br
slz.br
sorocaba.br
srv.br
taxi.br
teo.br
the.br
tmp.br
trd.br
tur.br
tv.br
udi.br
vet.br
vix.br
vlog.br
wiki.br
zlg.br

// bs : http://www.nic.bs/rules.html
bs
com.bs
net.bs
org.bs
edu.bs
gov.bs

// bt : https://en.wikipedia.org/wiki/.bt
bt
com.bt
edu.bt
gov.bt
net.bt
org.bt

// bv : No registrations at this time.
// Submitted by registry <jarle@uninett.no>
bv

// bw : https://en.wikipedia.org/wiki/.bw
// http://www.gobin.info/domainname/bw.doc
// list of other 2nd level tlds ?
bw
co.bw
org.bw

// by : https://en.wikipedia.org/wiki/.by
// http://tld.by/rules_2006_en.html
// list of other 2nd level tlds ?
by
gov.by
mil.by
// Official information does not indicate that com.by is a reserved
// second-level domain, but it's being used as one (see www.google.com.by and
// www.yahoo.com.by, for example), so we list it here for safety's sake.
com.by

// http://hoster.by/
of.by

// bz : https://en.wikipedia.org/wiki/.bz
// http://www.belizenic.bz/
bz
com.bz
net.bz
org.bz
edu.bz
gov.bz

// ca : https://en.wikipedia.org/wiki/.ca
ca
// ca geographical names
ab.ca
bc.ca
mb.ca
nb.ca
nf.ca
nl.ca
ns.ca
nt.ca
nu.ca
on.ca
pe.ca
qc.ca
sk.ca
yk.ca
// gc.ca: https://en.wikipedia.org/wiki/.gc.ca
// see also: http://registry.gc.ca/en/SubdomainFAQ
gc.ca

// cat : https://en.wikipedia.org/wiki/.cat
cat

// cc : https://en.wikipedia.org/wiki/.cc
cc

// cd : https://en.wikipedia.org/wiki/.cd
// see also: https://www.nic.cd/domain/insertDomain_2.jsp?act=1
cd
gov.cd

// cf : https://en.wikipedia.org/wiki/.cf
cf

// cg : https://en.wikipedia.org/wiki/.cg
cg

// ch : https://en.wikipedia.org/wiki/.ch
ch

// ci : https://en.wikipedia.org/wiki/.ci
// http://www.nic.ci/index.php?page=charte
ci
org.ci
or.ci
com.ci
co.ci
edu.ci
ed.ci
ac.ci
net.ci
go.ci
asso.ci
aéroport.ci
int.ci
presse.ci
md.ci
gouv.ci

// ck : https://en.wikipedia.org/wiki/.ck
*.ck
!www.ck

// cl : https://en.wikipedia.org/wiki/.cl
cl
gov.cl
gob.cl
co.cl
mil.cl

// cm : https://en.wikipedia.org/wiki/.cm plus bug 981927
cm
co.cm
com.cm
gov.cm
net.cm

// cn : https://en.wikipedia.org/wiki/.cn
// Submitted by registry <tanyaling@cnnic.cn>
cn
ac.cn
com.cn
edu.cn
gov.cn
net.cn
org.cn
mil.cn
公司.cn
网络.cn
網絡.cn
// cn geographic names
ah.cn
bj.cn
cq.cn
fj.cn
gd.cn
gs.cn
gz.cn
gx.cn
ha.cn
hb.cn
he.cn
hi.cn
hl.cn
hn.cn
jl.cn
js.cn
jx.cn
ln.cn
nm.cn
nx.cn
qh.cn
sc.cn
sd.cn
sh.cn
sn.cn
sx.cn
tj.cn
xj.cn
xz.cn
yn.cn
zj.cn
hk.cn
mo.cn
tw.cn

// co : https://en.wikipedia.org/wiki/.co
// Submitted by registry <tecnico@uniandes.edu.co>
co
arts.co
com.co
edu.co
firm.co
gov.co
info.co
int.co
mil.co
net.co
nom.co
org.co
rec.co
web.co

// com : https://en.wikipedia.org/wiki/.com
com

// coop : https://en.wikipedia.org/wiki/.coop
coop

// cr : http://www.nic.cr/niccr_publico/showRegistroDominiosScreen.do
cr
ac.cr
co.cr
ed.cr
fi.cr
go.cr
or.cr
sa.cr

// cu : https://en.wikipedia.org/wiki/.cu
cu
com.cu
edu.cu
org.cu
net.cu
gov.cu
inf.cu

// cv : https://en.wikipedia.org/wiki/.cv
cv

// cw : http://www.una.cw/cw_registry/
// Confirmed by registry <registry@una.net> 2013-03-26
cw
com.cw
edu.cw
net.cw
org.cw

// cx : https://en.wikipedia.org/wiki/.cx
// list of other 2nd level tlds ?
cx
gov.cx

// cy : http://www.nic.cy/
// Submitted by registry Panayiotou Fotia <cydns@ucy.ac.cy>
cy
ac.cy
biz.cy
com.cy
ekloges.cy
gov.cy
ltd.cy
name.cy
net.cy
org.cy
parliament.cy
press.cy
pro.cy
tm.cy

// cz : https://en.wikipedia.org/wiki/.cz
cz

// de : https://en.wikipedia.org/wiki/.de
// Confirmed by registry <ops@denic.de> (with technical
// reservations) 2008-07-01
de

// dj : https://en.wikipedia.org/wiki/.dj
dj

// dk : https://en.wikipedia.org/wiki/.dk
// Confirmed by registry <robert@dk-hostmaster.dk> 2008-06-17
dk

// dm : https://en.wikipedia.org/wiki/.dm
dm
com.dm
net.dm
org.dm
edu.dm
gov.dm

// do : https://en.wikipedia.org/wiki/.do
do
art.do
com.do
edu.do
gob.do
gov.do
mil.do
net.do
org.do
sld.do
web.do

// dz : https://en.wikipedia.org/wiki/.dz
dz
com.dz
org.dz
net.dz
gov.dz
edu.dz
asso.dz
pol.dz
art.dz

// ec : http://www.nic.ec/reg/paso1.asp
// Submitted by registry <vabboud@nic.ec>
ec
com.ec
info.ec
net.ec
fin.ec
k12.ec
med.ec
pro.ec
org.ec
edu.ec
gov.ec
gob.ec
mil.ec

// edu : https://en.wikipedia.org/wiki/.edu
edu

// ee : http://www.eenet.ee/EENet/dom_reeglid.html#lisa_B
ee
edu.ee
gov.ee
riik.ee
lib.ee
med.ee
com.ee
pri.ee
aip.ee
org.ee
fie.ee

// eg : https://en.wikipedia.org/wiki/.eg
eg
com.eg
edu.eg
eun.eg
gov.eg
mil.eg
name.eg
net.eg
org.eg
sci.eg

// er : https://en.wikipedia.org/wiki/.er
*.er

// es : https://www.nic.es/site_ingles/ingles/dominios/index.html
es
com.es
nom.es
org.es
gob.es
edu.es

// et : https://en.wikipedia.org/wiki/.et
et
com.et
gov.et
org.et
edu.et
biz.et
name.et
info.et
net.et

// eu : https://en.wikipedia.org/wiki/.eu
eu

// fi : https://en.wikipedia.org/wiki/.fi
fi
// aland.fi : https://en.wikipedia.org/wiki/.ax
// This domain is being phased out in favor of .ax. As there are still many
// domains under aland.fi, we still keep it on the list until aland.fi is
// completely removed.
// TODO: Check for updates (expected to be phased out around Q1/2009)
aland.fi

// fj : https://en.wikipedia.org/wiki/.fj
*.fj

// fk : https://en.wikipedia.org/wiki/.fk
*.fk

// fm : https://en.wikipedia.org/wiki/.fm
fm

// fo : https://en.wikipedia.org/wiki/.fo
fo

// fr : http://www.afnic.fr/
// domaines descriptifs : http://www.afnic.fr/obtenir/chartes/nommage-fr/annexe-descriptifs
fr
com.fr
asso.fr
nom.fr
prd.fr
presse.fr
tm.fr
// domaines sectoriels : http://www.afnic.fr/obtenir/chartes/nommage-fr/annexe-sectoriels
aeroport.fr
assedic.fr
avocat.fr
avoues.fr
cci.fr
chambagri.fr
chirurgiens-dentistes.fr
experts-comptables.fr
geometre-expert.fr
gouv.fr
greta.fr
huissier-justice.fr
medecin.fr
notaires.fr
pharmacien.fr
port.fr
veterinaire.fr

// ga : https://en.wikipedia.org/wiki/.ga
ga

// gb : This registry is effectively dormant
// Submitted by registry <Damien.Shaw@ja.net>
gb

// gd : https://en.wikipedia.org/wiki/.gd
gd

// ge : http://www.nic.net.ge/policy_en.pdf
ge
com.ge
edu.ge
gov.ge
org.ge
mil.ge
net.ge
pvt.ge

// gf : https://en.wikipedia.org/wiki/.gf
gf

// gg : http://www.channelisles.net/register-domains/
// Confirmed by registry <nigel@channelisles.net> 2013-11-28
gg
co.gg
net.gg
org.gg

// gh : https://en.wikipedia.org/wiki/.gh
// see also: http://www.nic.gh/reg_now.php
// Although domains directly at second level are not possible at the moment,
// they have been possible for some time and may come back.
gh
com.gh
edu.gh
gov.gh
org.gh
mil.gh

// gi : http://www.nic.gi/rules.html
gi
com.gi
ltd.gi
gov.gi
mod.gi
edu.gi
org.gi

// gl : https://en.wikipedia.org/wiki/.gl
// http://nic.gl
gl
co.gl
com.gl
edu.gl
net.gl
org.gl

// gm : http://www.nic.gm/htmlpages%5Cgm-policy.htm
gm

// gn : http://psg.com/dns/gn/gn.txt
// Submitted by registry <randy@psg.com>
gn
ac.gn
com.gn
edu.gn
gov.gn
org.gn
net.gn

// gov : https://en.wikipedia.org/wiki/.gov
gov

// gp : http://www.nic.gp/index.php?lang=en
gp
com.gp
net.gp
mobi.gp
edu.gp
org.gp
asso.gp

// gq : https://en.wikipedia.org/wiki/.gq
gq

// gr : https://grweb.ics.forth.gr/english/1617-B-2005.html
// Submitted by registry <segred@ics.forth.gr>
gr
com.gr
edu.gr
net.gr
org.gr
gov.gr

// gs : https://en.wikipedia.org/wiki/.gs
gs

// gt : http://www.gt/politicas_de_registro.html
gt
com.gt
edu.gt
gob.gt
ind.gt
mil.gt
net.gt
org.gt

// gu : http://gadao.gov.gu/register.html
// University of Guam : https://www.uog.edu
// Submitted by uognoc@triton.uog.edu
gu
com.gu
edu.gu
gov.gu
guam.gu
info.gu
net.gu
org.gu
web.gu

// gw : https://en.wikipedia.org/wiki/.gw
gw

// gy : https://en.wikipedia.org/wiki/.gy
// http://registry.gy/
gy
co.gy
com.gy
edu.gy
gov.gy
net.gy
org.gy

// hk : https://www.hkirc.hk
// Submitted by registry <hk.tech@hkirc.hk>
hk
com.hk
edu.hk
gov.hk
idv.hk
net.hk
org.hk
公司.hk
教育.hk
敎育.hk
政府.hk
個人.hk
个人.hk
箇人.hk
網络.hk
网络.hk
组織.hk
網絡.hk
网絡.hk
组织.hk
組織.hk
組织.hk

// hm : https://en.wikipedia.org/wiki/.hm
hm

// hn : http://www.nic.hn/politicas/ps02,,05.html
hn
com.hn
edu.hn
org.hn
net.hn
mil.hn
gob.hn

// hr : http://www.dns.hr/documents/pdf/HRTLD-regulations.pdf
hr
iz.hr
from.hr
name.hr
com.hr

// ht : http://www.nic.ht/info/charte.cfm
ht
com.ht
shop.ht
firm.ht
info.ht
adult.ht
net.ht
pro.ht
org.ht
med.ht
art.ht
coop.ht
pol.ht
asso.ht
edu.ht
rel.ht
gouv.ht
perso.ht

// hu : http://www.domain.hu/domain/English/sld.html
// Confirmed by registry <pasztor@iszt.hu> 2008-06-12
hu
co.hu
info.hu
org.hu
priv.hu
sport.hu
tm.hu
2000.hu
agrar.hu
bolt.hu
casino.hu
city.hu
erotica.hu
erotika.hu
film.hu
forum.hu
games.hu
hotel.hu
ingatlan.hu
jogasz.hu
konyvelo.hu
lakas.hu
media.hu
news.hu
reklam.hu
sex.hu
shop.hu
suli.hu
szex.hu
tozsde.hu
utazas.hu
video.hu

// id : https://pandi.id/en/domain/registration-requirements/
id
ac.id
biz.id
co.id
desa.id
go.id
mil.id
my.id
net.id
or.id
ponpes.id
sch.id
web.id

// ie : https://en.wikipedia.org/wiki/.ie
ie
gov.ie

// il : http://www.isoc.org.il/domains/
il
ac.il
co.il
gov.il
idf.il
k12.il
muni.il
net.il
org.il

// im : https://www.nic.im/
// Submitted by registry <info@nic.im>
im
ac.im
co.im
com.im
ltd.co.im
net.im
org.im
plc.co.im
tt.im
tv.im

// in : https://en.wikipedia.org/wiki/.in
// see also: https://registry.in/Policies
// Please note, that nic.in is not an official eTLD, but used by most
// government institutions.
in
co.in
firm.in
net.in
org.in
gen.in
ind.in
nic.in
ac.in
edu.in
res.in
gov.in
mil.in

// info : https://en.wikipedia.org/wiki/.info
info

// int : https://en.wikipedia.org/wiki/.int
// Confirmed by registry <iana-questions@icann.org> 2008-06-18
int
eu.int

// io : http://www.nic.io/rules.html
// list of other 2nd level tlds ?
io
com.io

// iq : http://www.cmc.iq/english/iq/iqregister1.htm
iq
gov.iq
edu.iq
mil.iq
com.iq
org.iq
net.iq

// ir : http://www.nic.ir/Terms_and_Conditions_ir,_Appendix_1_Domain_Rules
// Also see http://www.nic.ir/Internationalized_Domain_Names
// Two <iran>.ir entries added at request of <tech-team@nic.ir>, 2010-04-16
ir
ac.ir
co.ir
gov.ir
id.ir
net.ir
org.ir
sch.ir
// xn--mgba3a4f16a.ir (<iran>.ir, Persian YEH)
ایران.ir
// xn--mgba3a4fra.ir (<iran>.ir, Arabic YEH)
ايران.ir

// is : http://www.isnic.is/domain/rules.php
// Confirmed by registry <marius@isgate.is> 2008-12-06
is
net.is
com.is
edu.is
gov.is
org.is
int.is

// it : https://en.wikipedia.org/wiki/.it
it
gov.it
edu.it
// Reserved geo-names (regions and provinces):
// http://www.nic.it/sites/default/files/docs/Regulation_assignation_v7.1.pdf
// Regions
abr.it
abruzzo.it
aosta-valley.it
aostavalley.it
bas.it
basilicata.it
cal.it
calabria.it
cam.it
campania.it
emilia-romagna.it
emiliaromagna.it
emr.it
friuli-v-giulia.it
friuli-ve-giulia.it
friuli-vegiulia.it
friuli-venezia-giulia.it
friuli-veneziagiulia.it
friuli-vgiulia.it
friuliv-giulia.it
friulive-giulia.it
friulivegiulia.it
friulivenezia-giulia.it
friuliveneziagiulia.it
friulivgiulia.it
fvg.it
laz.it
lazio.it
lig.it
liguria.it
lom.it
lombardia.it
lombardy.it
lucania.it
mar.it
marche.it
mol.it
molise.it
piedmont.it
piemonte.it
pmn.it
pug.it
puglia.it
sar.it
sardegna.it
sardinia.it
sic.it
sicilia.it
sicily.it
taa.it
tos.it
toscana.it
trentin-sud-tirol.it
trentin-süd-tirol.it
trentin-sudtirol.it
trentin-südtirol.it
trentin-sued-tirol.it
trentin-suedtirol.it
trentino-a-adige.it
trentino-aadige.it
trentino-alto-adige.it
trentino-altoadige.it
trentino-s-tirol.it
trentino-stirol.it
trentino-sud-tirol.it
trentino-süd-tirol.it
trentino-sudtirol.it
trentino-südtirol.it
trentino-sued-tirol.it
trentino-suedtirol.it
trentino.it
trentinoa-adige.it
trentinoaadige.it
trentinoalto-adige.it
trentinoaltoadige.it
trentinos-tirol.it
trentinostirol.it
trentinosud-tirol.it
trentinosüd-tirol.it
trentinosudtirol.it
trentinosüdtirol.it
trentinosued-tirol.it
trentinosuedtirol.it
trentinsud-tirol.it
trentinsüd-tirol.it
trentinsudtirol.it
trentinsüdtirol.it
trentinsued-tirol.it
trentinsuedtirol.it
tuscany.it
umb.it
umbria.it
val-d-aosta.it
val-daosta.it
vald-aosta.it
valdaosta.it
valle-aosta.it
valle-d-aosta.it
valle-daosta.it
valleaosta.it
valled-aosta.it
valledaosta.it
vallee-aoste.it
vallée-aoste.it
vallee-d-aoste.it
vallée-d-aoste.it
valleeaoste.it
valléeaoste.it
valleedaoste.it
valléedaoste.it
vao.it
vda.it
ven.it
veneto.it
// Provinces
ag.it
agrigento.it
al.it
alessandria.it
alto-adige.it
altoadige.it
an.it
ancona.it
andria-barletta-trani.it
andria-trani-barletta.it
andriabarlettatrani.it
andriatranibarletta.it
ao.it
aosta.it
aoste.it
ap.it
aq.it
aquila.it
ar.it
arezzo.it
ascoli-piceno.it
ascolipiceno.it
asti.it
at.it
av.it
avellino.it
ba.it
balsan-sudtirol.it
balsan-südtirol.it
balsan-suedtirol.it
balsan.it
bari.it
barletta-trani-andria.it
barlettatraniandria.it
belluno.it
benevento.it
bergamo.it
bg.it
bi.it
biella.it
bl.it
bn.it
bo.it
bologna.it
bolzano-altoadige.it
bolzano.it
bozen-sudtirol.it
bozen-südtirol.it
bozen-suedtirol.it
bozen.it
br.it
brescia.it
brindisi.it
bs.it
bt.it
bulsan-sudtirol.it
bulsan-südtirol.it
bulsan-suedtirol.it
bulsan.it
bz.it
ca.it
cagliari.it
caltanissetta.it
campidano-medio.it
campidanomedio.it
campobasso.it
carbonia-iglesias.it
carboniaiglesias.it
carrara-massa.it
carraramassa.it
caserta.it
catania.it
catanzaro.it
cb.it
ce.it
cesena-forli.it
cesena-forlì.it
cesenaforli.it
cesenaforlì.it
ch.it
chieti.it
ci.it
cl.it
cn.it
co.it
como.it
cosenza.it
cr.it
cremona.it
crotone.it
cs.it
ct.it
cuneo.it
cz.it
dell-ogliastra.it
dellogliastra.it
en.it
enna.it
fc.it
fe.it
fermo.it
ferrara.it
fg.it
fi.it
firenze.it
florence.it
fm.it
foggia.it
forli-cesena.it
forlì-cesena.it
forlicesena.it
forlìcesena.it
fr.it
frosinone.it
ge.it
genoa.it
genova.it
go.it
gorizia.it
gr.it
grosseto.it
iglesias-carbonia.it
iglesiascarbonia.it
im.it
imperia.it
is.it
isernia.it
kr.it
la-spezia.it
laquila.it
laspezia.it
latina.it
lc.it
le.it
lecce.it
lecco.it
li.it
livorno.it
lo.it
lodi.it
lt.it
lu.it
lucca.it
macerata.it
mantova.it
massa-carrara.it
massacarrara.it
matera.it
mb.it
mc.it
me.it
medio-campidano.it
mediocampidano.it
messina.it
mi.it
milan.it
milano.it
mn.it
mo.it
modena.it
monza-brianza.it
monza-e-della-brianza.it
monza.it
monzabrianza.it
monzaebrianza.it
monzaedellabrianza.it
ms.it
mt.it
na.it
naples.it
napoli.it
no.it
novara.it
nu.it
nuoro.it
og.it
ogliastra.it
olbia-tempio.it
olbiatempio.it
or.it
oristano.it
ot.it
pa.it
padova.it
padua.it
palermo.it
parma.it
pavia.it
pc.it
pd.it
pe.it
perugia.it
pesaro-urbino.it
pesarourbino.it
pescara.it
pg.it
pi.it
piacenza.it
pisa.it
pistoia.it
pn.it
po.it
pordenone.it
potenza.it
pr.it
prato.it
pt.it
pu.it
pv.it
pz.it
ra.it
ragusa.it
ravenna.it
rc.it
re.it
reggio-calabria.it
reggio-emilia.it
reggiocalabria.it
reggioemilia.it
rg.it
ri.it
rieti.it
rimini.it
rm.it
rn.it
ro.it
roma.it
rome.it
rovigo.it
sa.it
salerno.it
sassari.it
savona.it
si.it
siena.it
siracusa.it
so.it
sondrio.it
sp.it
sr.it
ss.it
suedtirol.it
südtirol.it
sv.it
ta.it
taranto.it
te.it
tempio-olbia.it
tempioolbia.it
teramo.it
terni.it
tn.it
to.it
torino.it
tp.it
tr.it
trani-andria-barletta.it
trani-barletta-andria.it
traniandriabarletta.it
tranibarlettaandria.it
trapani.it
trento.it
treviso.it
trieste.it
ts.it
turin.it
tv.it
ud.it
udine.it
urbino-pesaro.it
urbinopesaro.it
va.it
varese.it
vb.it
vc.it
ve.it
venezia.it
venice.it
verbania.it
vercelli.it
verona.it
vi.it
vibo-valentia.it
vibovalentia.it
vicenza.it
viterbo.it
vr.it
vs.it
vt.it
vv.it

// je : http://www.channelisles.net/register-domains/
// Confirmed by registry <nigel@channelisles.net> 2013-11-28
je
co.je
net.je
org.je

// jm : http://www.com.jm/register.html
*.jm

// jo : http://www.dns.jo/Registration_policy.aspx
jo
com.jo
org.jo
net.jo
edu.jo
sch.jo
gov.jo
mil.jo
name.jo

// jobs : https://en.wikipedia.org/wiki/.jobs
jobs

// jp : https://en.wikipedia.org/wiki/.jp
// http://jprs.co.jp/en/jpdomain.html
// Submitted by registry <info@jprs.jp>
jp
// jp organizational type names
ac.jp
ad.jp
co.jp
ed.jp
go.jp
gr.jp
lg.jp
ne.jp
or.jp
// jp prefecture type names
aichi.jp
akita.jp
aomori.jp
chiba.jp
ehime.jp
fukui.jp
fukuoka.jp
fukushima.jp
gifu.jp
gunma.jp
hiroshima.jp
hokkaido.jp
hyogo.jp
ibaraki.jp
ishikawa.jp
iwate.jp
kagawa.jp
kagoshima.jp
kanagawa.jp
kochi.jp
kumamoto.jp
kyoto.jp
mie.jp
miyagi.jp
miyazaki.jp
nagano.jp
nagasaki.jp
nara.jp
niigata.jp
oita.jp
okayama.jp
okinawa.jp
osaka.jp
saga.jp
saitama.jp
shiga.jp
shimane.jp
shizuoka.jp
tochigi.jp
tokushima.jp
tokyo.jp
tottori.jp
toyama.jp
wakayama.jp
yamagata.jp
yamaguchi.jp
yamanashi.jp
栃木.jp
愛知.jp
愛媛.jp
兵庫.jp
熊本.jp
茨城.jp
北海道.jp
千葉.jp
和歌山.jp
長崎.jp
長野.jp
新潟.jp
青森.jp
静岡.jp
東京.jp
石川.jp
埼玉.jp
三重.jp
京都.jp
佐賀.jp
大分.jp
大阪.jp
奈良.jp
宮城.jp
宮崎.jp
富山.jp
山口.jp
山形.jp
山梨.jp
岩手.jp
岐阜.jp
岡山.jp
島根.jp
広島.jp
徳島.jp
沖縄.jp
滋賀.jp
神奈川.jp
福井.jp
福岡.jp
福島.jp
秋田.jp
群馬.jp
香川.jp
高知.jp
鳥取.jp
鹿児島.jp
// jp geographic type names
// http://jprs.jp/doc/rule/saisoku-1.html
*.kawasaki.jp
*.kitakyushu.jp
*.kobe.jp
*.nagoya.jp
*.sapporo.jp
*.sendai.jp
*.yokohama.jp
!city.kawasaki.jp
!city.kitakyushu.jp
!city.kobe.jp
!city.nagoya.jp
!city.sapporo.jp
!city.sendai.jp
!city.yokohama.jp
// 4th level registration
aisai.aichi.jp
ama.aichi.jp
anjo.aichi.jp
asuke.aichi.jp
chiryu.aichi.jp
chita.aichi.jp
fuso.aichi.jp
gamagori.aichi.jp
handa.aichi.jp
hazu.aichi.jp
hekinan.aichi.jp
higashiura.aichi.jp
ichinomiya.aichi.jp
inazawa.aichi.jp
inuyama.aichi.jp
isshiki.aichi.jp
iwakura.aichi.jp
kanie.aichi.jp
kariya.aichi.jp
kasugai.aichi.jp
kira.aichi.jp
kiyosu.aichi.jp
komaki.aichi.jp
konan.aichi.jp
kota.aichi.jp
mihama.aichi.jp
miyoshi.aichi.jp
nishio.aichi.jp
nisshin.aichi.jp
obu.aichi.jp
oguchi.aichi.jp
oharu.aichi.jp
okazaki.aichi.jp
owariasahi.aichi.jp
seto.aichi.jp
shikatsu.aichi.jp
shinshiro.aichi.jp
shitara.aichi.jp
tahara.aichi.jp
takahama.aichi.jp
tobishima.aichi.jp
toei.aichi.jp
togo.aichi.jp
tokai.aichi.jp
tokoname.aichi.jp
toyoake.aichi.jp
toyohashi.aichi.jp
toyokawa.aichi.jp
toyone.aichi.jp
toyota.aichi.jp
tsushima.aichi.jp
yatomi.aichi.jp
akita.akita.jp
daisen.akita.jp
fujisato.akita.jp
gojome.akita.jp
hachirogata.akita.jp
happou.akita.jp
higashinaruse.akita.jp
honjo.akita.jp
honjyo.akita.jp
ikawa.akita.jp
kamikoani.akita.jp
kamioka.akita.jp
katagami.akita.jp
kazuno.akita.jp
kitaakita.akita.jp
kosaka.akita.jp
kyowa.akita.jp
misato.akita.jp
mitane.akita.jp
moriyoshi.akita.jp
nikaho.akita.jp
noshiro.akita.jp
odate.akita.jp
oga.akita.jp
ogata.akita.jp
semboku.akita.jp
yokote.akita.jp
yurihonjo.akita.jp
aomori.aomori.jp
gonohe.aomori.jp
hachinohe.aomori.jp
hashikami.aomori.jp
hiranai.aomori.jp
hirosaki.aomori.jp
itayanagi.aomori.jp
kuroishi.aomori.jp
misawa.aomori.jp
mutsu.aomori.jp
nakadomari.aomori.jp
noheji.aomori.jp
oirase.aomori.jp
owani.aomori.jp
rokunohe.aomori.jp
sannohe.aomori.jp
shichinohe.aomori.jp
shingo.aomori.jp
takko.aomori.jp
towada.aomori.jp
tsugaru.aomori.jp
tsuruta.aomori.jp
abiko.chiba.jp
asahi.chiba.jp
chonan.chiba.jp
chosei.chiba.jp
choshi.chiba.jp
chuo.chiba.jp
funabashi.chiba.jp
futtsu.chiba.jp
hanamigawa.chiba.jp
ichihara.chiba.jp
ichikawa.chiba.jp
ichinomiya.chiba.jp
inzai.chiba.jp
isumi.chiba.jp
kamagaya.chiba.jp
kamogawa.chiba.jp
kashiwa.chiba.jp
katori.chiba.jp
katsuura.chiba.jp
kimitsu.chiba.jp
kisarazu.chiba.jp
kozaki.chiba.jp
kujukuri.chiba.jp
kyonan.chiba.jp
matsudo.chiba.jp
midori.chiba.jp
mihama.chiba.jp
minamiboso.chiba.jp
mobara.chiba.jp
mutsuzawa.chiba.jp
nagara.chiba.jp
nagareyama.chiba.jp
narashino.chiba.jp
narita.chiba.jp
noda.chiba.jp
oamishirasato.chiba.jp
omigawa.chiba.jp
onjuku.chiba.jp
otaki.chiba.jp
sakae.chiba.jp
sakura.chiba.jp
shimofusa.chiba.jp
shirako.chiba.jp
shiroi.chiba.jp
shisui.chiba.jp
sodegaura.chiba.jp
sosa.chiba.jp
tako.chiba.jp
tateyama.chiba.jp
togane.chiba.jp
tohnosho.chiba.jp
tomisato.chiba.jp
urayasu.chiba.jp
yachimata.chiba.jp
yachiyo.chiba.jp
yokaichiba.chiba.jp
yokoshibahikari.chiba.jp
yotsukaido.chiba.jp
ainan.ehime.jp
honai.ehime.jp
ikata.ehime.jp
imabari.ehime.jp
iyo.ehime.jp
kamijima.ehime.jp
kihoku.ehime.jp
kumakogen.ehime.jp
masaki.ehime.jp
matsuno.ehime.jp
matsuyama.ehime.jp
namikata.ehime.jp
niihama.ehime.jp
ozu.ehime.jp
saijo.ehime.jp
seiyo.ehime.jp
shikokuchuo.ehime.jp
tobe.ehime.jp
toon.ehime.jp
uchiko.ehime.jp
uwajima.ehime.jp
yawatahama.ehime.jp
echizen.fukui.jp
eiheiji.fukui.jp
fukui.fukui.jp
ikeda.fukui.jp
katsuyama.fukui.jp
mihama.fukui.jp
minamiechizen.fukui.jp
obama.fukui.jp
ohi.fukui.jp
ono.fukui.jp
sabae.fukui.jp
sakai.fukui.jp
takahama.fukui.jp
tsuruga.fukui.jp
wakasa.fukui.jp
ashiya.fukuoka.jp
buzen.fukuoka.jp
chikugo.fukuoka.jp
chikuho.fukuoka.jp
chikujo.fukuoka.jp
chikushino.fukuoka.jp
chikuzen.fukuoka.jp
chuo.fukuoka.jp
dazaifu.fukuoka.jp
fukuchi.fukuoka.jp
hakata.fukuoka.jp
higashi.fukuoka.jp
hirokawa.fukuoka.jp
hisayama.fukuoka.jp
iizuka.fukuoka.jp
inatsuki.fukuoka.jp
kaho.fukuoka.jp
kasuga.fukuoka.jp
kasuya.fukuoka.jp
kawara.fukuoka.jp
keisen.fukuoka.jp
koga.fukuoka.jp
kurate.fukuoka.jp
kurogi.fukuoka.jp
kurume.fukuoka.jp
minami.fukuoka.jp
miyako.fukuoka.jp
miyama.fukuoka.jp
miyawaka.fukuoka.jp
mizumaki.fukuoka.jp
munakata.fukuoka.jp
nakagawa.fukuoka.jp
nakama.fukuoka.jp
nishi.fukuoka.jp
nogata.fukuoka.jp
ogori.fukuoka.jp
okagaki.fukuoka.jp
okawa.fukuoka.jp
oki.fukuoka.jp
omuta.fukuoka.jp
onga.fukuoka.jp
onojo.fukuoka.jp
oto.fukuoka.jp
saigawa.fukuoka.jp
sasaguri.fukuoka.jp
shingu.fukuoka.jp
shinyoshitomi.fukuoka.jp
shonai.fukuoka.jp
soeda.fukuoka.jp
sue.fukuoka.jp
tachiarai.fukuoka.jp
tagawa.fukuoka.jp
takata.fukuoka.jp
toho.fukuoka.jp
toyotsu.fukuoka.jp
tsuiki.fukuoka.jp
ukiha.fukuoka.jp
umi.fukuoka.jp
usui.fukuoka.jp
yamada.fukuoka.jp
yame.fukuoka.jp
yanagawa.fukuoka.jp
yukuhashi.fukuoka.jp
aizubange.fukushima.jp
aizumisato.fukushima.jp
aizuwakamatsu.fukushima.jp
asakawa.fukushima.jp
bandai.fukushima.jp
date.fukushima.jp
fukushima.fukushima.jp
furudono.fukushima.jp
futaba.fukushima.jp
hanawa.fukushima.jp
higashi.fukushima.jp
hirata.fukushima.jp
hirono.fukushima.jp
iitate.fukushima.jp
inawashiro.fukushima.jp
ishikawa.fukushima.jp
iwaki.fukushima.jp
izumizaki.fukushima.jp
kagamiishi.fukushima.jp
kaneyama.fukushima.jp
kawamata.fukushima.jp
kitakata.fukushima.jp
kitashiobara.fukushima.jp
koori.fukushima.jp
koriyama.fukushima.jp
kunimi.fukushima.jp
miharu.fukushima.jp
mishima.fukushima.jp
namie.fukushima.jp
nango.fukushima.jp
nishiaizu.fukushima.jp
nishigo.fukushima.jp
okuma.fukushima.jp
omotego.fukushima.jp
ono.fukushima.jp
otama.fukushima.jp
samegawa.fukushima.jp
shimogo.fukushima.jp
shirakawa.fukushima.jp
showa.fukushima.jp
soma.fukushima.jp
sukagawa.fukushima.jp
taishin.fukushima.jp
tamakawa.fukushima.jp
tanagura.fukushima.jp
tenei.fukushima.jp
yabuki.fukushima.jp
yamato.fukushima.jp
yamatsuri.fukushima.jp
yanaizu.fukushima.jp
yugawa.fukushima.jp
anpachi.gifu.jp
ena.gifu.jp
gifu.gifu.jp
ginan.gifu.jp
godo.gifu.jp
gujo.gifu.jp
hashima.gifu.jp
hichiso.gifu.jp
hida.gifu.jp
higashishirakawa.gifu.jp
ibigawa.gifu.jp
ikeda.gifu.jp
kakamigahara.gifu.jp
kani.gifu.jp
kasahara.gifu.jp
kasamatsu.gifu.jp
kawaue.gifu.jp
kitagata.gifu.jp
mino.gifu.jp
minokamo.gifu.jp
mitake.gifu.jp
mizunami.gifu.jp
motosu.gifu.jp
nakatsugawa.gifu.jp
ogaki.gifu.jp
sakahogi.gifu.jp
seki.gifu.jp
sekigahara.gifu.jp
shirakawa.gifu.jp
tajimi.gifu.jp
takayama.gifu.jp
tarui.gifu.jp
toki.gifu.jp
tomika.gifu.jp
wanouchi.gifu.jp
yamagata.gifu.jp
yaotsu.gifu.jp
yoro.gifu.jp
annaka.gunma.jp
chiyoda.gunma.jp
fujioka.gunma.jp
higashiagatsuma.gunma.jp
isesaki.gunma.jp
itakura.gunma.jp
kanna.gunma.jp
kanra.gunma.jp
katashina.gunma.jp
kawaba.gunma.jp
kiryu.gunma.jp
kusatsu.gunma.jp
maebashi.gunma.jp
meiwa.gunma.jp
midori.gunma.jp
minakami.gunma.jp
naganohara.gunma.jp
nakanojo.gunma.jp
nanmoku.gunma.jp
numata.gunma.jp
oizumi.gunma.jp
ora.gunma.jp
ota.gunma.jp
shibukawa.gunma.jp
shimonita.gunma.jp
shinto.gunma.jp
showa.gunma.jp
takasaki.gunma.jp
takayama.gunma.jp
tamamura.gunma.jp
tatebayashi.gunma.jp
tomioka.gunma.jp
tsukiyono.gunma.jp
tsumagoi.gunma.jp
ueno.gunma.jp
yoshioka.gunma.jp
asaminami.hiroshima.jp
daiwa.hiroshima.jp
etajima.hiroshima.jp
fuchu.hiroshima.jp
fukuyama.hiroshima.jp
hatsukaichi.hiroshima.jp
higashihiroshima.hiroshima.jp
hongo.hiroshima.jp
jinsekikogen.hiroshima.jp
kaita.hiroshima.jp
kui.hiroshima.jp
kumano.hiroshima.jp
kure.hiroshima.jp
mihara.hiroshima.jp
miyoshi.hiroshima.jp
naka.hiroshima.jp
onomichi.hiroshima.jp
osakikamijima.hiroshima.jp
otake.hiroshima.jp
saka.hiroshima.jp
sera.hiroshima.jp
seranishi.hiroshima.jp
shinichi.hiroshima.jp
shobara.hiroshima.jp
takehara.hiroshima.jp
abashiri.hokkaido.jp
abira.hokkaido.jp
aibetsu.hokkaido.jp
akabira.hokkaido.jp
akkeshi.hokkaido.jp
asahikawa.hokkaido.jp
ashibetsu.hokkaido.jp
ashoro.hokkaido.jp
assabu.hokkaido.jp
atsuma.hokkaido.jp
bibai.hokkaido.jp
biei.hokkaido.jp
bifuka.hokkaido.jp
bihoro.hokkaido.jp
biratori.hokkaido.jp
chippubetsu.hokkaido.jp
chitose.hokkaido.jp
date.hokkaido.jp
ebetsu.hokkaido.jp
embetsu.hokkaido.jp
eniwa.hokkaido.jp
erimo.hokkaido.jp
esan.hokkaido.jp
esashi.hokkaido.jp
fukagawa.hokkaido.jp
fukushima.hokkaido.jp
furano.hokkaido.jp
furubira.hokkaido.jp
haboro.hokkaido.jp
hakodate.hokkaido.jp
hamatonbetsu.hokkaido.jp
hidaka.hokkaido.jp
higashikagura.hokkaido.jp
higashikawa.hokkaido.jp
hiroo.hokkaido.jp
hokuryu.hokkaido.jp
hokuto.hokkaido.jp
honbetsu.hokkaido.jp
horokanai.hokkaido.jp
horonobe.hokkaido.jp
ikeda.hokkaido.jp
imakane.hokkaido.jp
ishikari.hokkaido.jp
iwamizawa.hokkaido.jp
iwanai.hokkaido.jp
kamifurano.hokkaido.jp
kamikawa.hokkaido.jp
kamishihoro.hokkaido.jp
kamisunagawa.hokkaido.jp
kamoenai.hokkaido.jp
kayabe.hokkaido.jp
kembuchi.hokkaido.jp
kikonai.hokkaido.jp
kimobetsu.hokkaido.jp
kitahiroshima.hokkaido.jp
kitami.hokkaido.jp
kiyosato.hokkaido.jp
koshimizu.hokkaido.jp
kunneppu.hokkaido.jp
kuriyama.hokkaido.jp
kuromatsunai.hokkaido.jp
kushiro.hokkaido.jp
kutchan.hokkaido.jp
kyowa.hokkaido.jp
mashike.hokkaido.jp
matsumae.hokkaido.jp
mikasa.hokkaido.jp
minamifurano.hokkaido.jp
mombetsu.hokkaido.jp
moseushi.hokkaido.jp
mukawa.hokkaido.jp
muroran.hokkaido.jp
naie.hokkaido.jp
nakagawa.hokkaido.jp
nakasatsunai.hokkaido.jp
nakatombetsu.hokkaido.jp
nanae.hokkaido.jp
nanporo.hokkaido.jp
nayoro.hokkaido.jp
nemuro.hokkaido.jp
niikappu.hokkaido.jp
niki.hokkaido.jp
nishiokoppe.hokkaido.jp
noboribetsu.hokkaido.jp
numata.hokkaido.jp
obihiro.hokkaido.jp
obira.hokkaido.jp
oketo.hokkaido.jp
okoppe.hokkaido.jp
otaru.hokkaido.jp
otobe.hokkaido.jp
otofuke.hokkaido.jp
otoineppu.hokkaido.jp
oumu.hokkaido.jp
ozora.hokkaido.jp
pippu.hokkaido.jp
rankoshi.hokkaido.jp
rebun.hokkaido.jp
rikubetsu.hokkaido.jp
rishiri.hokkaido.jp
rishirifuji.hokkaido.jp
saroma.hokkaido.jp
sarufutsu.hokkaido.jp
shakotan.hokkaido.jp
shari.hokkaido.jp
shibecha.hokkaido.jp
shibetsu.hokkaido.jp
shikabe.hokkaido.jp
shikaoi.hokkaido.jp
shimamaki.hokkaido.jp
shimizu.hokkaido.jp
shimokawa.hokkaido.jp
shinshinotsu.hokkaido.jp
shintoku.hokkaido.jp
shiranuka.hokkaido.jp
shiraoi.hokkaido.jp
shiriuchi.hokkaido.jp
sobetsu.hokkaido.jp
sunagawa.hokkaido.jp
taiki.hokkaido.jp
takasu.hokkaido.jp
takikawa.hokkaido.jp
takinoue.hokkaido.jp
teshikaga.hokkaido.jp
tobetsu.hokkaido.jp
tohma.hokkaido.jp
tomakomai.hokkaido.jp
tomari.hokkaido.jp
toya.hokkaido.jp
toyako.hokkaido.jp
toyotomi.hokkaido.jp
toyoura.hokkaido.jp
tsubetsu.hokkaido.jp
tsukigata.hokkaido.jp
urakawa.hokkaido.jp
urausu.hokkaido.jp
uryu.hokkaido.jp
utashinai.hokkaido.jp
wakkanai.hokkaido.jp
wassamu.hokkaido.jp
yakumo.hokkaido.jp
yoichi.hokkaido.jp
aioi.hyogo.jp
akashi.hyogo.jp
ako.hyogo.jp
amagasaki.hyogo.jp
aogaki.hyogo.jp
asago.hyogo.jp
ashiya.hyogo.jp
awaji.hyogo.jp
fukusaki.hyogo.jp
goshiki.hyogo.jp
harima.hyogo.jp
himeji.hyogo.jp
ichikawa.hyogo.jp
inagawa.hyogo.jp
itami.hyogo.jp
kakogawa.hyogo.jp
kamigori.hyogo.jp
kamikawa.hyogo.jp
kasai.hyogo.jp
kasuga.hyogo.jp
kawanishi.hyogo.jp
miki.hyogo.jp
minamiawaji.hyogo.jp
nishinomiya.hyogo.jp
nishiwaki.hyogo.jp
ono.hyogo.jp
sanda.hyogo.jp
sannan.hyogo.jp
sasayama.hyogo.jp
sayo.hyogo.jp
shingu.hyogo.jp
shinonsen.hyogo.jp
shiso.hyogo.jp
sumoto.hyogo.jp
taishi.hyogo.jp
taka.hyogo.jp
takarazuka.hyogo.jp
takasago.hyogo.jp
takino.hyogo.jp
tamba.hyogo.jp
tatsuno.hyogo.jp
toyooka.hyogo.jp
yabu.hyogo.jp
yashiro.hyogo.jp
yoka.hyogo.jp
yokawa.hyogo.jp
ami.ibaraki.jp
asahi.ibaraki.jp
bando.ibaraki.jp
chikusei.ibaraki.jp
daigo.ibaraki.jp
fujishiro.ibaraki.jp
hitachi.ibaraki.jp
hitachinaka.ibaraki.jp
hitachiomiya.ibaraki.jp
hitachiota.ibaraki.jp
ibaraki.ibaraki.jp
ina.ibaraki.jp
inashiki.ibaraki.jp
itako.ibaraki.jp
iwama.ibaraki.jp
joso.ibaraki.jp
kamisu.ibaraki.jp
kasama.ibaraki.jp
kashima.ibaraki.jp
kasumigaura.ibaraki.jp
koga.ibaraki.jp
miho.ibaraki.jp
mito.ibaraki.jp
moriya.ibaraki.jp
naka.ibaraki.jp
namegata.ibaraki.jp
oarai.ibaraki.jp
ogawa.ibaraki.jp
omitama.ibaraki.jp
ryugasaki.ibaraki.jp
sakai.ibaraki.jp
sakuragawa.ibaraki.jp
shimodate.ibaraki.jp
shimotsuma.ibaraki.jp
shirosato.ibaraki.jp
sowa.ibaraki.jp
suifu.ibaraki.jp
takahagi.ibaraki.jp
tamatsukuri.ibaraki.jp
tokai.ibaraki.jp
tomobe.ibaraki.jp
tone.ibaraki.jp
toride.ibaraki.jp
tsuchiura.ibaraki.jp
tsukuba.ibaraki.jp
uchihara.ibaraki.jp
ushiku.ibaraki.jp
yachiyo.ibaraki.jp
yamagata.ibaraki.jp
yawara.ibaraki.jp
yuki.ibaraki.jp
anamizu.ishikawa.jp
hakui.ishikawa.jp
hakusan.ishikawa.jp
kaga.ishikawa.jp
kahoku.ishikawa.jp
kanazawa.ishikawa.jp
kawakita.ishikawa.jp
komatsu.ishikawa.jp
nakanoto.ishikawa.jp
nanao.ishikawa.jp
nomi.ishikawa.jp
nonoichi.ishikawa.jp
noto.ishikawa.jp
shika.ishikawa.jp
suzu.ishikawa.jp
tsubata.ishikawa.jp
tsurugi.ishikawa.jp
uchinada.ishikawa.jp
wajima.ishikawa.jp
fudai.iwate.jp
fujisawa.iwate.jp
hanamaki.iwate.jp
hiraizumi.iwate.jp
hirono.iwate.jp
ichinohe.iwate.jp
ichinoseki.iwate.jp
iwaizumi.iwate.jp
iwate.iwate.jp
joboji.iwate.jp
kamaishi.iwate.jp
kanegasaki.iwate.jp
karumai.iwate.jp
kawai.iwate.jp
kitakami.iwate.jp
kuji.iwate.jp
kunohe.iwate.jp
kuzumaki.iwate.jp
miyako.iwate.jp
mizusawa.iwate.jp
morioka.iwate.jp
ninohe.iwate.jp
noda.iwate.jp
ofunato.iwate.jp
oshu.iwate.jp
otsuchi.iwate.jp
rikuzentakata.iwate.jp
shiwa.iwate.jp
shizukuishi.iwate.jp
sumita.iwate.jp
tanohata.iwate.jp
tono.iwate.jp
yahaba.iwate.jp
yamada.iwate.jp
ayagawa.kagawa.jp
higashikagawa.kagawa.jp
kanonji.kagawa.jp
kotohira.kagawa.jp
manno.kagawa.jp
marugame.kagawa.jp
mitoyo.kagawa.jp
naoshima.kagawa.jp
sanuki.kagawa.jp
tadotsu.kagawa.jp
takamatsu.kagawa.jp
tonosho.kagawa.jp
uchinomi.kagawa.jp
utazu.kagawa.jp
zentsuji.kagawa.jp
akune.kagoshima.jp
amami.kagoshima.jp
hioki.kagoshima.jp
isa.kagoshima.jp
isen.kagoshima.jp
izumi.kagoshima.jp
kagoshima.kagoshima.jp
kanoya.kagoshima.jp
kawanabe.kagoshima.jp
kinko.kagoshima.jp
kouyama.kagoshima.jp
makurazaki.kagoshima.jp
matsumoto.kagoshima.jp
minamitane.kagoshima.jp
nakatane.kagoshima.jp
nishinoomote.kagoshima.jp
satsumasendai.kagoshima.jp
soo.kagoshima.jp
tarumizu.kagoshima.jp
yusui.kagoshima.jp
aikawa.kanagawa.jp
atsugi.kanagawa.jp
ayase.kanagawa.jp
chigasaki.kanagawa.jp
ebina.kanagawa.jp
fujisawa.kanagawa.jp
hadano.kanagawa.jp
hakone.kanagawa.jp
hiratsuka.kanagawa.jp
isehara.kanagawa.jp
kaisei.kanagawa.jp
kamakura.kanagawa.jp
kiyokawa.kanagawa.jp
matsuda.kanagawa.jp
minamiashigara.kanagawa.jp
miura.kanagawa.jp
nakai.kanagawa.jp
ninomiya.kanagawa.jp
odawara.kanagawa.jp
oi.kanagawa.jp
oiso.kanagawa.jp
sagamihara.kanagawa.jp
samukawa.kanagawa.jp
tsukui.kanagawa.jp
yamakita.kanagawa.jp
yamato.kanagawa.jp
yokosuka.kanagawa.jp
yugawara.kanagawa.jp
zama.kanagawa.jp
zushi.kanagawa.jp
aki.kochi.jp
geisei.kochi.jp
hidaka.kochi.jp
higashitsuno.kochi.jp
ino.kochi.jp
kagami.kochi.jp
kami.kochi.jp
kitagawa.kochi.jp
kochi.kochi.jp
mihara.kochi.jp
motoyama.kochi.jp
muroto.kochi.jp
nahari.kochi.jp
nakamura.kochi.jp
nankoku.kochi.jp
nishitosa.kochi.jp
niyodogawa.kochi.jp
ochi.kochi.jp
okawa.kochi.jp
otoyo.kochi.jp
otsuki.kochi.jp
sakawa.kochi.jp
sukumo.kochi.jp
susaki.kochi.jp
tosa.kochi.jp
tosashimizu.kochi.jp
toyo.kochi.jp
tsuno.kochi.jp
umaji.kochi.jp
yasuda.kochi.jp
yusuhara.kochi.jp
amakusa.kumamoto.jp
arao.kumamoto.jp
aso.kumamoto.jp
choyo.kumamoto.jp
gyokuto.kumamoto.jp
kamiamakusa.kumamoto.jp
kikuchi.kumamoto.jp
kumamoto.kumamoto.jp
mashiki.kumamoto.jp
mifune.kumamoto.jp
minamata.kumamoto.jp
minamioguni.kumamoto.jp
nagasu.kumamoto.jp
nishihara.kumamoto.jp
oguni.kumamoto.jp
ozu.kumamoto.jp
sumoto.kumamoto.jp
takamori.kumamoto.jp
uki.kumamoto.jp
uto.kumamoto.jp
yamaga.kumamoto.jp
yamato.kumamoto.jp
yatsushiro.kumamoto.jp
ayabe.kyoto.jp
fukuchiyama.kyoto.jp
higashiyama.kyoto.jp
ide.kyoto.jp
ine.kyoto.jp
joyo.kyoto.jp
kameoka.kyoto.jp
kamo.kyoto.jp
kita.kyoto.jp
kizu.kyoto.jp
kumiyama.kyoto.jp
kyotamba.kyoto.jp
kyotanabe.kyoto.jp
kyotango.kyoto.jp
maizuru.kyoto.jp
minami.kyoto.jp
minamiyamashiro.kyoto.jp
miyazu.kyoto.jp
muko.kyoto.jp
nagaokakyo.kyoto.jp
nakagyo.kyoto.jp
nantan.kyoto.jp
oyamazaki.kyoto.jp
sakyo.kyoto.jp
seika.kyoto.jp
tanabe.kyoto.jp
uji.kyoto.jp
ujitawara.kyoto.jp
wazuka.kyoto.jp
yamashina.kyoto.jp
yawata.kyoto.jp
asahi.mie.jp
inabe.mie.jp
ise.mie.jp
kameyama.mie.jp
kawagoe.mie.jp
kiho.mie.jp
kisosaki.mie.jp
kiwa.mie.jp
komono.mie.jp
kumano.mie.jp
kuwana.mie.jp
matsusaka.mie.jp
meiwa.mie.jp
mihama.mie.jp
minamiise.mie.jp
misugi.mie.jp
miyama.mie.jp
nabari.mie.jp
shima.mie.jp
suzuka.mie.jp
tado.mie.jp
taiki.mie.jp
taki.mie.jp
tamaki.mie.jp
toba.mie.jp
tsu.mie.jp
udono.mie.jp
ureshino.mie.jp
watarai.mie.jp
yokkaichi.mie.jp
furukawa.miyagi.jp
higashimatsushima.miyagi.jp
ishinomaki.miyagi.jp
iwanuma.miyagi.jp
kakuda.miyagi.jp
kami.miyagi.jp
kawasaki.miyagi.jp
marumori.miyagi.jp
matsushima.miyagi.jp
minamisanriku.miyagi.jp
misato.miyagi.jp
murata.miyagi.jp
natori.miyagi.jp
ogawara.miyagi.jp
ohira.miyagi.jp
onagawa.miyagi.jp
osaki.miyagi.jp
rifu.miyagi.jp
semine.miyagi.jp
shibata.miyagi.jp
shichikashuku.miyagi.jp
shikama.miyagi.jp
shiogama.miyagi.jp
shiroishi.miyagi.jp
tagajo.miyagi.jp
taiwa.miyagi.jp
tome.miyagi.jp
tomiya.miyagi.jp
wakuya.miyagi.jp
watari.miyagi.jp
yamamoto.miyagi.jp
zao.miyagi.jp
aya.miyazaki.jp
ebino.miyazaki.jp
gokase.miyazaki.jp
hyuga.miyazaki.jp
kadogawa.miyazaki.jp
kawaminami.miyazaki.jp
kijo.miyazaki.jp
kitagawa.miyazaki.jp
kitakata.miyazaki.jp
kitaura.miyazaki.jp
kobayashi.miyazaki.jp
kunitomi.miyazaki.jp
kushima.miyazaki.jp
mimata.miyazaki.jp
miyakonojo.miyazaki.jp
miyazaki.miyazaki.jp
morotsuka.miyazaki.jp
nichinan.miyazaki.jp
nishimera.miyazaki.jp
nobeoka.miyazaki.jp
saito.miyazaki.jp
shiiba.miyazaki.jp
shintomi.miyazaki.jp
takaharu.miyazaki.jp
takanabe.miyazaki.jp
takazaki.miyazaki.jp
tsuno.miyazaki.jp
achi.nagano.jp
agematsu.nagano.jp
anan.nagano.jp
aoki.nagano.jp
asahi.nagano.jp
azumino.nagano.jp
chikuhoku.nagano.jp
chikuma.nagano.jp
chino.nagano.jp
fujimi.nagano.jp
hakuba.nagano.jp
hara.nagano.jp
hiraya.nagano.jp
iida.nagano.jp
iijima.nagano.jp
iiyama.nagano.jp
iizuna.nagano.jp
ikeda.nagano.jp
ikusaka.nagano.jp
ina.nagano.jp
karuizawa.nagano.jp
kawakami.nagano.jp
kiso.nagano.jp
kisofukushima.nagano.jp
kitaaiki.nagano.jp
komagane.nagano.jp
komoro.nagano.jp
matsukawa.nagano.jp
matsumoto.nagano.jp
miasa.nagano.jp
minamiaiki.nagano.jp
minamimaki.nagano.jp
minamiminowa.nagano.jp
minowa.nagano.jp
miyada.nagano.jp
miyota.nagano.jp
mochizuki.nagano.jp
nagano.nagano.jp
nagawa.nagano.jp
nagiso.nagano.jp
nakagawa.nagano.jp
nakano.nagano.jp
nozawaonsen.nagano.jp
obuse.nagano.jp
ogawa.nagano.jp
okaya.nagano.jp
omachi.nagano.jp
omi.nagano.jp
ookuwa.nagano.jp
ooshika.nagano.jp
otaki.nagano.jp
otari.nagano.jp
sakae.nagano.jp
sakaki.nagano.jp
saku.nagano.jp
sakuho.nagano.jp
shimosuwa.nagano.jp
shinanomachi.nagano.jp
shiojiri.nagano.jp
suwa.nagano.jp
suzaka.nagano.jp
takagi.nagano.jp
takamori.nagano.jp
takayama.nagano.jp
tateshina.nagano.jp
tatsuno.nagano.jp
togakushi.nagano.jp
togura.nagano.jp
tomi.nagano.jp
ueda.nagano.jp
wada.nagano.jp
yamagata.nagano.jp
yamanouchi.nagano.jp
yasaka.nagano.jp
yasuoka.nagano.jp
chijiwa.nagasaki.jp
futsu.nagasaki.jp
goto.nagasaki.jp
hasami.nagasaki.jp
hirado.nagasaki.jp
iki.nagasaki.jp
isahaya.nagasaki.jp
kawatana.nagasaki.jp
kuchinotsu.nagasaki.jp
matsuura.nagasaki.jp
nagasaki.nagasaki.jp
obama.nagasaki.jp
omura.nagasaki.jp
oseto.nagasaki.jp
saikai.nagasaki.jp
sasebo.nagasaki.jp
seihi.nagasaki.jp
shimabara.nagasaki.jp
shinkamigoto.nagasaki.jp
togitsu.nagasaki.jp
tsushima.nagasaki.jp
unzen.nagasaki.jp
ando.nara.jp
gose.nara.jp
heguri.nara.jp
higashiyoshino.nara.jp
ikaruga.nara.jp
ikoma.nara.jp
kamikitayama.nara.jp
kanmaki.nara.jp
kashiba.nara.jp
kashihara.nara.jp
katsuragi.nara.jp
kawai.nara.jp
kawakami.nara.jp
kawanishi.nara.jp
koryo.nara.jp
kurotaki.nara.jp
mitsue.nara.jp
miyake.nara.jp
nara.nara.jp
nosegawa.nara.jp
oji.nara.jp
ouda.nara.jp
oyodo.nara.jp
sakurai.nara.jp
sango.nara.jp
shimoichi.nara.jp
shimokitayama.nara.jp
shinjo.nara.jp
soni.nara.jp
takatori.nara.jp
tawaramoto.nara.jp
tenkawa.nara.jp
tenri.nara.jp
uda.nara.jp
yamatokoriyama.nara.jp
yamatotakada.nara.jp
yamazoe.nara.jp
yoshino.nara.jp
aga.niigata.jp
agano.niigata.jp
gosen.niigata.jp
itoigawa.niigata.jp
izumozaki.niigata.jp
joetsu.niigata.jp
kamo.niigata.jp
kariwa.niigata.jp
kashiwazaki.niigata.jp
minamiuonuma.niigata.jp
mitsuke.niigata.jp
muika.niigata.jp
murakami.niigata.jp
myoko.niigata.jp
nagaoka.niigata.jp
niigata.niigata.jp
ojiya.niigata.jp
omi.niigata.jp
sado.niigata.jp
sanjo.niigata.jp
seiro.niigata.jp
seirou.niigata.jp
sekikawa.niigata.jp
shibata.niigata.jp
tagami.niigata.jp
tainai.niigata.jp
tochio.niigata.jp
tokamachi.niigata.jp
tsubame.niigata.jp
tsunan.niigata.jp
uonuma.niigata.jp
yahiko.niigata.jp
yoita.niigata.jp
yuzawa.niigata.jp
beppu.oita.jp
bungoono.oita.jp
bungotakada.oita.jp
hasama.oita.jp
hiji.oita.jp
himeshima.oita.jp
hita.oita.jp
kamitsue.oita.jp
kokonoe.oita.jp
kuju.oita.jp
kunisaki.oita.jp
kusu.oita.jp
oita.oita.jp
saiki.oita.jp
taketa.oita.jp
tsukumi.oita.jp
usa.oita.jp
usuki.oita.jp
yufu.oita.jp
akaiwa.okayama.jp
asakuchi.okayama.jp
bizen.okayama.jp
hayashima.okayama.jp
ibara.okayama.jp
kagamino.okayama.jp
kasaoka.okayama.jp
kibichuo.okayama.jp
kumenan.okayama.jp
kurashiki.okayama.jp
maniwa.okayama.jp
misaki.okayama.jp
nagi.okayama.jp
niimi.okayama.jp
nishiawakura.okayama.jp
okayama.okayama.jp
satosho.okayama.jp
setouchi.okayama.jp
shinjo.okayama.jp
shoo.okayama.jp
soja.okayama.jp
takahashi.okayama.jp
tamano.okayama.jp
tsuyama.okayama.jp
wake.okayama.jp
yakage.okayama.jp
aguni.okinawa.jp
ginowan.okinawa.jp
ginoza.okinawa.jp
gushikami.okinawa.jp
haebaru.okinawa.jp
higashi.okinawa.jp
hirara.okinawa.jp
iheya.okinawa.jp
ishigaki.okinawa.jp
ishikawa.okinawa.jp
itoman.okinawa.jp
izena.okinawa.jp
kadena.okinawa.jp
kin.okinawa.jp
kitadaito.okinawa.jp
kitanakagusuku.okinawa.jp
kumejima.okinawa.jp
kunigami.okinawa.jp
minamidaito.okinawa.jp
motobu.okinawa.jp
nago.okinawa.jp
naha.okinawa.jp
nakagusuku.okinawa.jp
nakijin.okinawa.jp
nanjo.okinawa.jp
nishihara.okinawa.jp
ogimi.okinawa.jp
okinawa.okinawa.jp
onna.okinawa.jp
shimoji.okinawa.jp
taketomi.okinawa.jp
tarama.okinawa.jp
tokashiki.okinawa.jp
tomigusuku.okinawa.jp
tonaki.okinawa.jp
urasoe.okinawa.jp
uruma.okinawa.jp
yaese.okinawa.jp
yomitan.okinawa.jp
yonabaru.okinawa.jp
yonaguni.okinawa.jp
zamami.okinawa.jp
abeno.osaka.jp
chihayaakasaka.osaka.jp
chuo.osaka.jp
daito.osaka.jp
fujiidera.osaka.jp
habikino.osaka.jp
hannan.osaka.jp
higashiosaka.osaka.jp
higashisumiyoshi.osaka.jp
higashiyodogawa.osaka.jp
hirakata.osaka.jp
ibaraki.osaka.jp
ikeda.osaka.jp
izumi.osaka.jp
izumiotsu.osaka.jp
izumisano.osaka.jp
kadoma.osaka.jp
kaizuka.osaka.jp
kanan.osaka.jp
kashiwara.osaka.jp
katano.osaka.jp
kawachinagano.osaka.jp
kishiwada.osaka.jp
kita.osaka.jp
kumatori.osaka.jp
matsubara.osaka.jp
minato.osaka.jp
minoh.osaka.jp
misaki.osaka.jp
moriguchi.osaka.jp
neyagawa.osaka.jp
nishi.osaka.jp
nose.osaka.jp
osakasayama.osaka.jp
sakai.osaka.jp
sayama.osaka.jp
sennan.osaka.jp
settsu.osaka.jp
shijonawate.osaka.jp
shimamoto.osaka.jp
suita.osaka.jp
tadaoka.osaka.jp
taishi.osaka.jp
tajiri.osaka.jp
takaishi.osaka.jp
takatsuki.osaka.jp
tondabayashi.osaka.jp
toyonaka.osaka.jp
toyono.osaka.jp
yao.osaka.jp
ariake.saga.jp
arita.saga.jp
fukudomi.saga.jp
genkai.saga.jp
hamatama.saga.jp
hizen.saga.jp
imari.saga.jp
kamimine.saga.jp
kanzaki.saga.jp
karatsu.saga.jp
kashima.saga.jp
kitagata.saga.jp
kitahata.saga.jp
kiyama.saga.jp
kouhoku.saga.jp
kyuragi.saga.jp
nishiarita.saga.jp
ogi.saga.jp
omachi.saga.jp
ouchi.saga.jp
saga.saga.jp
shiroishi.saga.jp
taku.saga.jp
tara.saga.jp
tosu.saga.jp
yoshinogari.saga.jp
arakawa.saitama.jp
asaka.saitama.jp
chichibu.saitama.jp
fujimi.saitama.jp
fujimino.saitama.jp
fukaya.saitama.jp
hanno.saitama.jp
hanyu.saitama.jp
hasuda.saitama.jp
hatogaya.saitama.jp
hatoyama.saitama.jp
hidaka.saitama.jp
higashichichibu.saitama.jp
higashimatsuyama.saitama.jp
honjo.saitama.jp
ina.saitama.jp
iruma.saitama.jp
iwatsuki.saitama.jp
kamiizumi.saitama.jp
kamikawa.saitama.jp
kamisato.saitama.jp
kasukabe.saitama.jp
kawagoe.saitama.jp
kawaguchi.saitama.jp
kawajima.saitama.jp
kazo.saitama.jp
kitamoto.saitama.jp
koshigaya.saitama.jp
kounosu.saitama.jp
kuki.saitama.jp
kumagaya.saitama.jp
matsubushi.saitama.jp
minano.saitama.jp
misato.saitama.jp
miyashiro.saitama.jp
miyoshi.saitama.jp
moroyama.saitama.jp
nagatoro.saitama.jp
namegawa.saitama.jp
niiza.saitama.jp
ogano.saitama.jp
ogawa.saitama.jp
ogose.saitama.jp
okegawa.saitama.jp
omiya.saitama.jp
otaki.saitama.jp
ranzan.saitama.jp
ryokami.saitama.jp
saitama.saitama.jp
sakado.saitama.jp
satte.saitama.jp
sayama.saitama.jp
shiki.saitama.jp
shiraoka.saitama.jp
soka.saitama.jp
sugito.saitama.jp
toda.saitama.jp
tokigawa.saitama.jp
tokorozawa.saitama.jp
tsurugashima.saitama.jp
urawa.saitama.jp
warabi.saitama.jp
yashio.saitama.jp
yokoze.saitama.jp
yono.saitama.jp
yorii.saitama.jp
yoshida.saitama.jp
yoshikawa.saitama.jp
yoshimi.saitama.jp
aisho.shiga.jp
gamo.shiga.jp
higashiomi.shiga.jp
hikone.shiga.jp
koka.shiga.jp
konan.shiga.jp
kosei.shiga.jp
koto.shiga.jp
kusatsu.shiga.jp
maibara.shiga.jp
moriyama.shiga.jp
nagahama.shiga.jp
nishiazai.shiga.jp
notogawa.shiga.jp
omihachiman.shiga.jp
otsu.shiga.jp
ritto.shiga.jp
ryuoh.shiga.jp
takashima.shiga.jp
takatsuki.shiga.jp
torahime.shiga.jp
toyosato.shiga.jp
yasu.shiga.jp
akagi.shimane.jp
ama.shimane.jp
gotsu.shimane.jp
hamada.shimane.jp
higashiizumo.shimane.jp
hikawa.shimane.jp
hikimi.shimane.jp
izumo.shimane.jp
kakinoki.shimane.jp
masuda.shimane.jp
matsue.shimane.jp
misato.shimane.jp
nishinoshima.shimane.jp
ohda.shimane.jp
okinoshima.shimane.jp
okuizumo.shimane.jp
shimane.shimane.jp
tamayu.shimane.jp
tsuwano.shimane.jp
unnan.shimane.jp
yakumo.shimane.jp
yasugi.shimane.jp
yatsuka.shimane.jp
arai.shizuoka.jp
atami.shizuoka.jp
fuji.shizuoka.jp
fujieda.shizuoka.jp
fujikawa.shizuoka.jp
fujinomiya.shizuoka.jp
fukuroi.shizuoka.jp
gotemba.shizuoka.jp
haibara.shizuoka.jp
hamamatsu.shizuoka.jp
higashiizu.shizuoka.jp
ito.shizuoka.jp
iwata.shizuoka.jp
izu.shizuoka.jp
izunokuni.shizuoka.jp
kakegawa.shizuoka.jp
kannami.shizuoka.jp
kawanehon.shizuoka.jp
kawazu.shizuoka.jp
kikugawa.shizuoka.jp
kosai.shizuoka.jp
makinohara.shizuoka.jp
matsuzaki.shizuoka.jp
minamiizu.shizuoka.jp
mishima.shizuoka.jp
morimachi.shizuoka.jp
nishiizu.shizuoka.jp
numazu.shizuoka.jp
omaezaki.shizuoka.jp
shimada.shizuoka.jp
shimizu.shizuoka.jp
shimoda.shizuoka.jp
shizuoka.shizuoka.jp
susono.shizuoka.jp
yaizu.shizuoka.jp
yoshida.shizuoka.jp
ashikaga.tochigi.jp
bato.tochigi.jp
haga.tochigi.jp
ichikai.tochigi.jp
iwafune.tochigi.jp
kaminokawa.tochigi.jp
kanuma.tochigi.jp
karasuyama.tochigi.jp
kuroiso.tochigi.jp
mashiko.tochigi.jp
mibu.tochigi.jp
moka.tochigi.jp
motegi.tochigi.jp
nasu.tochigi.jp
nasushiobara.tochigi.jp
nikko.tochigi.jp
nishikata.tochigi.jp
nogi.tochigi.jp
ohira.tochigi.jp
ohtawara.tochigi.jp
oyama.tochigi.jp
sakura.tochigi.jp
sano.tochigi.jp
shimotsuke.tochigi.jp
shioya.tochigi.jp
takanezawa.tochigi.jp
tochigi.tochigi.jp
tsuga.tochigi.jp
ujiie.tochigi.jp
utsunomiya.tochigi.jp
yaita.tochigi.jp
aizumi.tokushima.jp
anan.tokushima.jp
ichiba.tokushima.jp
itano.tokushima.jp
kainan.tokushima.jp
komatsushima.tokushima.jp
matsushige.tokushima.jp
mima.tokushima.jp
minami.tokushima.jp
miyoshi.tokushima.jp
mugi.tokushima.jp
nakagawa.tokushima.jp
naruto.tokushima.jp
sanagochi.tokushima.jp
shishikui.tokushima.jp
tokushima.tokushima.jp
wajiki.tokushima.jp
adachi.tokyo.jp
akiruno.tokyo.jp
akishima.tokyo.jp
aogashima.tokyo.jp
arakawa.tokyo.jp
bunkyo.tokyo.jp
chiyoda.tokyo.jp
chofu.tokyo.jp
chuo.tokyo.jp
edogawa.tokyo.jp
fuchu.tokyo.jp
fussa.tokyo.jp
hachijo.tokyo.jp
hachioji.tokyo.jp
hamura.tokyo.jp
higashikurume.tokyo.jp
higashimurayama.tokyo.jp
higashiyamato.tokyo.jp
hino.tokyo.jp
hinode.tokyo.jp
hinohara.tokyo.jp
inagi.tokyo.jp
itabashi.tokyo.jp
katsushika.tokyo.jp
kita.tokyo.jp
kiyose.tokyo.jp
kodaira.tokyo.jp
koganei.tokyo.jp
kokubunji.tokyo.jp
komae.tokyo.jp
koto.tokyo.jp
kouzushima.tokyo.jp
kunitachi.tokyo.jp
machida.tokyo.jp
meguro.tokyo.jp
minato.tokyo.jp
mitaka.tokyo.jp
mizuho.tokyo.jp
musashimurayama.tokyo.jp
musashino.tokyo.jp
nakano.tokyo.jp
nerima.tokyo.jp
ogasawara.tokyo.jp
okutama.tokyo.jp
ome.tokyo.jp
oshima.tokyo.jp
ota.tokyo.jp
setagaya.tokyo.jp
shibuya.tokyo.jp
shinagawa.tokyo.jp
shinjuku.tokyo.jp
suginami.tokyo.jp
sumida.tokyo.jp
tachikawa.tokyo.jp
taito.tokyo.jp
tama.tokyo.jp
toshima.tokyo.jp
chizu.tottori.jp
hino.tottori.jp
kawahara.tottori.jp
koge.tottori.jp
kotoura.tottori.jp
misasa.tottori.jp
nanbu.tottori.jp
nichinan.tottori.jp
sakaiminato.tottori.jp
tottori.tottori.jp
wakasa.tottori.jp
yazu.tottori.jp
yonago.tottori.jp
asahi.toyama.jp
fuchu.toyama.jp
fukumitsu.toyama.jp
funahashi.toyama.jp
himi.toyama.jp
imizu.toyama.jp
inami.toyama.jp
johana.toyama.jp
kamiichi.toyama.jp
kurobe.toyama.jp
nakaniikawa.toyama.jp
namerikawa.toyama.jp
nanto.toyama.jp
nyuzen.toyama.jp
oyabe.toyama.jp
taira.toyama.jp
takaoka.toyama.jp
tateyama.toyama.jp
toga.toyama.jp
tonami.toyama.jp
toyama.toyama.jp
unazuki.toyama.jp
uozu.toyama.jp
yamada.toyama.jp
arida.wakayama.jp
aridagawa.wakayama.jp
gobo.wakayama.jp
hashimoto.wakayama.jp
hidaka.wakayama.jp
hirogawa.wakayama.jp
inami.wakayama.jp
iwade.wakayama.jp
kainan.wakayama.jp
kamitonda.wakayama.jp
katsuragi.wakayama.jp
kimino.wakayama.jp
kinokawa.wakayama.jp
kitayama.wakayama.jp
koya.wakayama.jp
koza.wakayama.jp
kozagawa.wakayama.jp
kudoyama.wakayama.jp
kushimoto.wakayama.jp
mihama.wakayama.jp
misato.wakayama.jp
nachikatsuura.wakayama.jp
shingu.wakayama.jp
shirahama.wakayama.jp
taiji.wakayama.jp
tanabe.wakayama.jp
wakayama.wakayama.jp
yuasa.wakayama.jp
yura.wakayama.jp
asahi.yamagata.jp
funagata.yamagata.jp
higashine.yamagata.jp
iide.yamagata.jp
kahoku.yamagata.jp
kaminoyama.yamagata.jp
kaneyama.yamagata.jp
kawanishi.yamagata.jp
mamurogawa.yamagata.jp
mikawa.yamagata.jp
murayama.yamagata.jp
nagai.yamagata.jp
nakayama.yamagata.jp
nanyo.yamagata.jp
nishikawa.yamagata.jp
obanazawa.yamagata.jp
oe.yamagata.jp
oguni.yamagata.jp
ohkura.yamagata.jp
oishida.yamagata.jp
sagae.yamagata.jp
sakata.yamagata.jp
sakegawa.yamagata.jp
shinjo.yamagata.jp
shirataka.yamagata.jp
shonai.yamagata.jp
takahata.yamagata.jp
tendo.yamagata.jp
tozawa.yamagata.jp
tsuruoka.yamagata.jp
yamagata.yamagata.jp
yamanobe.yamagata.jp
yonezawa.yamagata.jp
yuza.yamagata.jp
abu.yamaguchi.jp
hagi.yamaguchi.jp
hikari.yamaguchi.jp
hofu.yamaguchi.jp
iwakuni.yamaguchi.jp
kudamatsu.yamaguchi.jp
mitou.yamaguchi.jp
nagato.yamaguchi.jp
oshima.yamaguchi.jp
shimonoseki.yamaguchi.jp
shunan.yamaguchi.jp
tabuse.yamaguchi.jp
tokuyama.yamaguchi.jp
toyota.yamaguchi.jp
ube.yamaguchi.jp
yuu.yamaguchi.jp
chuo.yamanashi.jp
doshi.yamanashi.jp
fuefuki.yamanashi.jp
fujikawa.yamanashi.jp
fujikawaguchiko.yamanashi.jp
fujiyoshida.yamanashi.jp
hayakawa.yamanashi.jp
hokuto.yamanashi.jp
ichikawamisato.yamanashi.jp
kai.yamanashi.jp
kofu.yamanashi.jp
koshu.yamanashi.jp
kosuge.yamanashi.jp
minami-alps.yamanashi.jp
minobu.yamanashi.jp
nakamichi.yamanashi.jp
nanbu.yamanashi.jp
narusawa.yamanashi.jp
nirasaki.yamanashi.jp
nishikatsura.yamanashi.jp
oshino.yamanashi.jp
otsuki.yamanashi.jp
showa.yamanashi.jp
tabayama.yamanashi.jp
tsuru.yamanashi.jp
uenohara.yamanashi.jp
yamanakako.yamanashi.jp
yamanashi.yamanashi.jp

// ke : http://www.kenic.or.ke/index.php/en/ke-domains/ke-domains
ke
ac.ke
co.ke
go.ke
info.ke
me.ke
mobi.ke
ne.ke
or.ke
sc.ke

// kg : http://www.domain.kg/dmn_n.html
kg
org.kg
net.kg
com.kg
edu.kg
gov.kg
mil.kg

// kh : http://www.mptc.gov.kh/dns_registration.htm
*.kh

// ki : http://www.ki/dns/index.html
ki
edu.ki
biz.ki
net.ki
org.ki
gov.ki
info.ki
com.ki

// km : https://en.wikipedia.org/wiki/.km
// http://www.domaine.km/documents/charte.doc
km
org.km
nom.km
gov.km
prd.km
tm.km
edu.km
mil.km
ass.km
com.km
// These are only mentioned as proposed suggestions at domaine.km, but
// https://en.wikipedia.org/wiki/.km says they're available for registration:
coop.km
asso.km
presse.km
medecin.km
notaires.km
pharmaciens.km
veterinaire.km
gouv.km

// kn : https://en.wikipedia.org/wiki/.kn
// http://www.dot.kn/domainRules.html
kn
net.kn
org.kn
edu.kn
gov.kn

// kp : http://www.kcce.kp/en_index.php
kp
com.kp
edu.kp
gov.kp
org.kp
rep.kp
tra.kp

// kr : https://en.wikipedia.org/wiki/.kr
// see also: http://domain.nida.or.kr/eng/registration.jsp
kr
ac.kr
co.kr
es.kr
go.kr
hs.kr
kg.kr
mil.kr
ms.kr
ne.kr
or.kr
pe.kr
re.kr
sc.kr
// kr geographical names
busan.kr
chungbuk.kr
chungnam.kr
daegu.kr
daejeon.kr
gangwon.kr
gwangju.kr
gyeongbuk.kr
gyeonggi.kr
gyeongnam.kr
incheon.kr
jeju.kr
jeonbuk.kr
jeonnam.kr
seoul.kr
ulsan.kr

// kw : https://www.nic.kw/policies/
// Confirmed by registry <nic.tech@citra.gov.kw>
kw
com.kw
edu.kw
emb.kw
gov.kw
ind.kw
net.kw
org.kw

// ky : http://www.icta.ky/da_ky_reg_dom.php
// Confirmed by registry <kysupport@perimeterusa.com> 2008-06-17
ky
edu.ky
gov.ky
com.ky
org.ky
net.ky

// kz : https://en.wikipedia.org/wiki/.kz
// see also: http://www.nic.kz/rules/index.jsp
kz
org.kz
edu.kz
net.kz
gov.kz
mil.kz
com.kz

// la : https://en.wikipedia.org/wiki/.la
// Submitted by registry <gavin.brown@nic.la>
la
int.la
net.la
info.la
edu.la
gov.la
per.la
com.la
org.la

// lb : https://en.wikipedia.org/wiki/.lb
// Submitted by registry <randy@psg.com>
lb
com.lb
edu.lb
gov.lb
net.lb
org.lb

// lc : https://en.wikipedia.org/wiki/.lc
// see also: http://www.nic.lc/rules.htm
lc
com.lc
net.lc
co.lc
org.lc
edu.lc
gov.lc

// li : https://en.wikipedia.org/wiki/.li
li

// lk : http://www.nic.lk/seclevpr.html
lk
gov.lk
sch.lk
net.lk
int.lk
com.lk
org.lk
edu.lk
ngo.lk
soc.lk
web.lk
ltd.lk
assn.lk
grp.lk
hotel.lk
ac.lk

// lr : http://psg.com/dns/lr/lr.txt
// Submitted by registry <randy@psg.com>
lr
com.lr
edu.lr
gov.lr
org.lr
net.lr

// ls : https://en.wikipedia.org/wiki/.ls
ls
co.ls
org.ls

// lt : https://en.wikipedia.org/wiki/.lt
lt
// gov.lt : http://www.gov.lt/index_en.php
gov.lt

// lu : http://www.dns.lu/en/
lu

// lv : http://www.nic.lv/DNS/En/generic.php
lv
com.lv
edu.lv
gov.lv
org.lv
mil.lv
id.lv
net.lv
asn.lv
conf.lv

// ly : http://www.nic.ly/regulations.php
ly
com.ly
net.ly
gov.ly
plc.ly
edu.ly
sch.ly
med.ly
org.ly
id.ly

// ma : https://en.wikipedia.org/wiki/.ma
// http://www.anrt.ma/fr/admin/download/upload/file_fr782.pdf
ma
co.ma
net.ma
gov.ma
org.ma
ac.ma
press.ma

// mc : http://www.nic.mc/
mc
tm.mc
asso.mc

// md : https://en.wikipedia.org/wiki/.md
md

// me : https://en.wikipedia.org/wiki/.me
me
co.me
net.me
org.me
edu.me
ac.me
gov.me
its.me
priv.me

// mg : http://nic.mg/nicmg/?page_id=39
mg
org.mg
nom.mg
gov.mg
prd.mg
tm.mg
edu.mg
mil.mg
com.mg
co.mg

// mh : https://en.wikipedia.org/wiki/.mh
mh

// mil : https://en.wikipedia.org/wiki/.mil
mil

// mk : https://en.wikipedia.org/wiki/.mk
// see also: http://dns.marnet.net.mk/postapka.php
mk
com.mk
org.mk
net.mk
edu.mk
gov.mk
inf.mk
name.mk

// ml : http://www.gobin.info/domainname/ml-template.doc
// see also: https://en.wikipedia.org/wiki/.ml
ml
com.ml
edu.ml
gouv.ml
gov.ml
net.ml
org.ml
presse.ml

// mm : https://en.wikipedia.org/wiki/.mm
*.mm

// mn : https://en.wikipedia.org/wiki/.mn
mn
gov.mn
edu.mn
org.mn

// mo : http://www.monic.net.mo/
mo
com.mo
net.mo
org.mo
edu.mo
gov.mo

// mobi : https://en.wikipedia.org/wiki/.mobi
mobi

// mp : http://www.dot.mp/
// Confirmed by registry <dcamacho@saipan.com> 2008-06-17
mp

// mq : https://en.wikipedia.org/wiki/.mq
mq

// mr : https://en.wikipedia.org/wiki/.mr
mr
gov.mr

// ms : http://www.nic.ms/pdf/MS_Domain_Name_Rules.pdf
ms
com.ms
edu.ms
gov.ms
net.ms
org.ms

// mt : https://www.nic.org.mt/go/policy
// Submitted by registry <help@nic.org.mt>
mt
com.mt
edu.mt
net.mt
org.mt

// mu : https://en.wikipedia.org/wiki/.mu
mu
com.mu
net.mu
org.mu
gov.mu
ac.mu
co.mu
or.mu

// museum : http://about.museum/naming/
// http://index.museum/
museum
academy.museum
agriculture.museum
air.museum
airguard.museum
alabama.museum
alaska.museum
amber.museum
ambulance.museum
american.museum
americana.museum
americanantiques.museum
americanart.museum
amsterdam.museum
and.museum
annefrank.museum
anthro.museum
anthropology.museum
antiques.museum
aquarium.museum
arboretum.museum
archaeological.museum
archaeology.museum
architecture.museum
art.museum
artanddesign.museum
artcenter.museum
artdeco.museum
arteducation.museum
artgallery.museum
arts.museum
artsandcrafts.museum
asmatart.museum
assassination.museum
assisi.museum
association.museum
astronomy.museum
atlanta.museum
austin.museum
australia.museum
automotive.museum
aviation.museum
axis.museum
badajoz.museum
baghdad.museum
bahn.museum
bale.museum
baltimore.museum
barcelona.museum
baseball.museum
basel.museum
baths.museum
bauern.museum
beauxarts.museum
beeldengeluid.museum
bellevue.museum
bergbau.museum
berkeley.museum
berlin.museum
bern.museum
bible.museum
bilbao.museum
bill.museum
birdart.museum
birthplace.museum
bonn.museum
boston.museum
botanical.museum
botanicalgarden.museum
botanicgarden.museum
botany.museum
brandywinevalley.museum
brasil.museum
bristol.museum
british.museum
britishcolumbia.museum
broadcast.museum
brunel.museum
brussel.museum
brussels.museum
bruxelles.museum
building.museum
burghof.museum
bus.museum
bushey.museum
cadaques.museum
california.museum
cambridge.museum
can.museum
canada.museum
capebreton.museum
carrier.museum
cartoonart.museum
casadelamoneda.museum
castle.museum
castres.museum
celtic.museum
center.museum
chattanooga.museum
cheltenham.museum
chesapeakebay.museum
chicago.museum
children.museum
childrens.museum
childrensgarden.museum
chiropractic.museum
chocolate.museum
christiansburg.museum
cincinnati.museum
cinema.museum
circus.museum
civilisation.museum
civilization.museum
civilwar.museum
clinton.museum
clock.museum
coal.museum
coastaldefence.museum
cody.museum
coldwar.museum
collection.museum
colonialwilliamsburg.museum
coloradoplateau.museum
columbia.museum
columbus.museum
communication.museum
communications.museum
community.museum
computer.museum
computerhistory.museum
comunicações.museum
contemporary.museum
contemporaryart.museum
convent.museum
copenhagen.museum
corporation.museum
correios-e-telecomunicações.museum
corvette.museum
costume.museum
countryestate.museum
county.museum
crafts.museum
cranbrook.museum
creation.museum
cultural.museum
culturalcenter.museum
culture.museum
cyber.museum
cymru.museum
dali.museum
dallas.museum
database.museum
ddr.museum
decorativearts.museum
delaware.museum
delmenhorst.museum
denmark.museum
depot.museum
design.museum
detroit.museum
dinosaur.museum
discovery.museum
dolls.museum
donostia.museum
durham.museum
eastafrica.museum
eastcoast.museum
education.museum
educational.museum
egyptian.museum
eisenbahn.museum
elburg.museum
elvendrell.museum
embroidery.museum
encyclopedic.museum
england.museum
entomology.museum
environment.museum
environmentalconservation.museum
epilepsy.museum
essex.museum
estate.museum
ethnology.museum
exeter.museum
exhibition.museum
family.museum
farm.museum
farmequipment.museum
farmers.museum
farmstead.museum
field.museum
figueres.museum
filatelia.museum
film.museum
fineart.museum
finearts.museum
finland.museum
flanders.museum
florida.museum
force.museum
fortmissoula.museum
fortworth.museum
foundation.museum
francaise.museum
frankfurt.museum
franziskaner.museum
freemasonry.museum
freiburg.museum
fribourg.museum
frog.museum
fundacio.museum
furniture.museum
gallery.museum
garden.museum
gateway.museum
geelvinck.museum
gemological.museum
geology.museum
georgia.museum
giessen.museum
glas.museum
glass.museum
gorge.museum
grandrapids.museum
graz.museum
guernsey.museum
halloffame.museum
hamburg.museum
handson.museum
harvestcelebration.museum
hawaii.museum
health.museum
heimatunduhren.museum
hellas.museum
helsinki.museum
hembygdsforbund.museum
heritage.museum
histoire.museum
historical.museum
historicalsociety.museum
historichouses.museum
historisch.museum
historisches.museum
history.museum
historyofscience.museum
horology.museum
house.museum
humanities.museum
illustration.museum
imageandsound.museum
indian.museum
indiana.museum
indianapolis.museum
indianmarket.museum
intelligence.museum
interactive.museum
iraq.museum
iron.museum
isleofman.museum
jamison.museum
jefferson.museum
jerusalem.museum
jewelry.museum
jewish.museum
jewishart.museum
jfk.museum
journalism.museum
judaica.museum
judygarland.museum
juedisches.museum
juif.museum
karate.museum
karikatur.museum
kids.museum
koebenhavn.museum
koeln.museum
kunst.museum
kunstsammlung.museum
kunstunddesign.museum
labor.museum
labour.museum
lajolla.museum
lancashire.museum
landes.museum
lans.museum
läns.museum
larsson.museum
lewismiller.museum
lincoln.museum
linz.museum
living.museum
livinghistory.museum
localhistory.museum
london.museum
losangeles.museum
louvre.museum
loyalist.museum
lucerne.museum
luxembourg.museum
luzern.museum
mad.museum
madrid.museum
mallorca.museum
manchester.museum
mansion.museum
mansions.museum
manx.museum
marburg.museum
maritime.museum
maritimo.museum
maryland.museum
marylhurst.museum
media.museum
medical.museum
medizinhistorisches.museum
meeres.museum
memorial.museum
mesaverde.museum
michigan.museum
midatlantic.museum
military.museum
mill.museum
miners.museum
mining.museum
minnesota.museum
missile.museum
missoula.museum
modern.museum
moma.museum
money.museum
monmouth.museum
monticello.museum
montreal.museum
moscow.museum
motorcycle.museum
muenchen.museum
muenster.museum
mulhouse.museum
muncie.museum
museet.museum
museumcenter.museum
museumvereniging.museum
music.museum
national.museum
nationalfirearms.museum
nationalheritage.museum
nativeamerican.museum
naturalhistory.museum
naturalhistorymuseum.museum
naturalsciences.museum
nature.museum
naturhistorisches.museum
natuurwetenschappen.museum
naumburg.museum
naval.museum
nebraska.museum
neues.museum
newhampshire.museum
newjersey.museum
newmexico.museum
newport.museum
newspaper.museum
newyork.museum
niepce.museum
norfolk.museum
north.museum
nrw.museum
nuernberg.museum
nuremberg.museum
nyc.museum
nyny.museum
oceanographic.museum
oceanographique.museum
omaha.museum
online.museum
ontario.museum
openair.museum
oregon.museum
oregontrail.museum
otago.museum
oxford.museum
pacific.museum
paderborn.museum
palace.museum
paleo.museum
palmsprings.museum
panama.museum
paris.museum
pasadena.museum
pharmacy.museum
philadelphia.museum
philadelphiaarea.museum
philately.museum
phoenix.museum
photography.museum
pilots.museum
pittsburgh.museum
planetarium.museum
plantation.museum
plants.museum
plaza.museum
portal.museum
portland.museum
portlligat.museum
posts-and-telecommunications.museum
preservation.museum
presidio.museum
press.museum
project.museum
public.museum
pubol.museum
quebec.museum
railroad.museum
railway.museum
research.museum
resistance.museum
riodejaneiro.museum
rochester.museum
rockart.museum
roma.museum
russia.museum
saintlouis.museum
salem.museum
salvadordali.museum
salzburg.museum
sandiego.museum
sanfrancisco.museum
santabarbara.museum
santacruz.museum
santafe.museum
saskatchewan.museum
satx.museum
savannahga.museum
schlesisches.museum
schoenbrunn.museum
schokoladen.museum
school.museum
schweiz.museum
science.museum
scienceandhistory.museum
scienceandindustry.museum
sciencecenter.museum
sciencecenters.museum
science-fiction.museum
sciencehistory.museum
sciences.museum
sciencesnaturelles.museum
scotland.museum
seaport.museum
settlement.museum
settlers.museum
shell.museum
sherbrooke.museum
sibenik.museum
silk.museum
ski.museum
skole.museum
society.museum
sologne.museum
soundandvision.museum
southcarolina.museum
southwest.museum
space.museum
spy.museum
square.museum
stadt.museum
stalbans.museum
starnberg.museum
state.museum
stateofdelaware.museum
station.museum
steam.museum
steiermark.museum
stjohn.museum
stockholm.museum
stpetersburg.museum
stuttgart.museum
suisse.museum
surgeonshall.museum
surrey.museum
svizzera.museum
sweden.museum
sydney.museum
tank.museum
tcm.museum
technology.museum
telekommunikation.museum
television.museum
texas.museum
textile.museum
theater.museum
time.museum
timekeeping.museum
topology.museum
torino.museum
touch.museum
town.museum
transport.museum
tree.museum
trolley.museum
trust.museum
trustee.museum
uhren.museum
ulm.museum
undersea.museum
university.museum
usa.museum
usantiques.museum
usarts.museum
uscountryestate.museum
usculture.museum
usdecorativearts.museum
usgarden.museum
ushistory.museum
ushuaia.museum
uslivinghistory.museum
utah.museum
uvic.museum
valley.museum
vantaa.museum
versailles.museum
viking.museum
village.museum
virginia.museum
virtual.museum
virtuel.museum
vlaanderen.museum
volkenkunde.museum
wales.museum
wallonie.museum
war.museum
washingtondc.museum
watchandclock.museum
watch-and-clock.museum
western.museum
westfalen.museum
whaling.museum
wildlife.museum
williamsburg.museum
windmill.museum
workshop.museum
york.museum
yorkshire.museum
yosemite.museum
youth.museum
zoological.museum
zoology.museum
ירושלים.museum
иком.museum

// mv : https://en.wikipedia.org/wiki/.mv
// "mv" included because, contra Wikipedia, google.mv exists.
mv
aero.mv
biz.mv
com.mv
coop.mv
edu.mv
gov.mv
info.mv
int.mv
mil.mv
museum.mv
name.mv
net.mv
org.mv
pro.mv

// mw : http://www.registrar.mw/
mw
ac.mw
biz.mw
co.mw
com.mw
coop.mw
edu.mw
gov.mw
int.mw
museum.mw
net.mw
org.mw

// mx : http://www.nic.mx/
// Submitted by registry <farias@nic.mx>
mx
com.mx
org.mx
gob.mx
edu.mx
net.mx

// my : http://www.mynic.net.my/
my
com.my
net.my
org.my
gov.my
edu.my
mil.my
name.my

// mz : http://www.uem.mz/
// Submitted by registry <antonio@uem.mz>
mz
ac.mz
adv.mz
co.mz
edu.mz
gov.mz
mil.mz
net.mz
org.mz

// na : http://www.na-nic.com.na/
// http://www.info.na/domain/
na
info.na
pro.na
name.na
school.na
or.na
dr.na
us.na
mx.na
ca.na
in.na
cc.na
tv.na
ws.na
mobi.na
co.na
com.na
org.na

// name : has 2nd-level tlds, but there's no list of them
name

// nc : http://www.cctld.nc/
nc
asso.nc
nom.nc

// ne : https://en.wikipedia.org/wiki/.ne
ne

// net : https://en.wikipedia.org/wiki/.net
net

// nf : https://en.wikipedia.org/wiki/.nf
nf
com.nf
net.nf
per.nf
rec.nf
web.nf
arts.nf
firm.nf
info.nf
other.nf
store.nf

// ng : http://www.nira.org.ng/index.php/join-us/register-ng-domain/189-nira-slds
ng
com.ng
edu.ng
gov.ng
i.ng
mil.ng
mobi.ng
name.ng
net.ng
org.ng
sch.ng

// ni : http://www.nic.ni/
ni
ac.ni
biz.ni
co.ni
com.ni
edu.ni
gob.ni
in.ni
info.ni
int.ni
mil.ni
net.ni
nom.ni
org.ni
web.ni

// nl : https://en.wikipedia.org/wiki/.nl
//      https://www.sidn.nl/
//      ccTLD for the Netherlands
nl

// BV.nl will be a registry for dutch BV's (besloten vennootschap)
bv.nl

// no : http://www.norid.no/regelverk/index.en.html
// The Norwegian registry has declined to notify us of updates. The web pages
// referenced below are the official source of the data. There is also an
// announce mailing list:
// https://postlister.uninett.no/sympa/info/norid-diskusjon
no
// Norid generic domains : http://www.norid.no/regelverk/vedlegg-c.en.html
fhs.no
vgs.no
fylkesbibl.no
folkebibl.no
museum.no
idrett.no
priv.no
// Non-Norid generic domains : http://www.norid.no/regelverk/vedlegg-d.en.html
mil.no
stat.no
dep.no
kommune.no
herad.no
// no geographical names : http://www.norid.no/regelverk/vedlegg-b.en.html
// counties
aa.no
ah.no
bu.no
fm.no
hl.no
hm.no
jan-mayen.no
mr.no
nl.no
nt.no
of.no
ol.no
oslo.no
rl.no
sf.no
st.no
svalbard.no
tm.no
tr.no
va.no
vf.no
// primary and lower secondary schools per county
gs.aa.no
gs.ah.no
gs.bu.no
gs.fm.no
gs.hl.no
gs.hm.no
gs.jan-mayen.no
gs.mr.no
gs.nl.no
gs.nt.no
gs.of.no
gs.ol.no
gs.oslo.no
gs.rl.no
gs.sf.no
gs.st.no
gs.svalbard.no
gs.tm.no
gs.tr.no
gs.va.no
gs.vf.no
// cities
akrehamn.no
åkrehamn.no
algard.no
ålgård.no
arna.no
brumunddal.no
bryne.no
bronnoysund.no
brønnøysund.no
drobak.no
drøbak.no
egersund.no
fetsund.no
floro.no
florø.no
fredrikstad.no
hokksund.no
honefoss.no
hønefoss.no
jessheim.no
jorpeland.no
jørpeland.no
kirkenes.no
kopervik.no
krokstadelva.no
langevag.no
langevåg.no
leirvik.no
mjondalen.no
mjøndalen.no
mo-i-rana.no
mosjoen.no
mosjøen.no
nesoddtangen.no
orkanger.no
osoyro.no
osøyro.no
raholt.no
råholt.no
sandnessjoen.no
sandnessjøen.no
skedsmokorset.no
slattum.no
spjelkavik.no
stathelle.no
stavern.no
stjordalshalsen.no
stjørdalshalsen.no
tananger.no
tranby.no
vossevangen.no
// communities
afjord.no
åfjord.no
agdenes.no
al.no
ål.no
alesund.no
ålesund.no
alstahaug.no
alta.no
áltá.no
alaheadju.no
álaheadju.no
alvdal.no
amli.no
åmli.no
amot.no
åmot.no
andebu.no
andoy.no
andøy.no
andasuolo.no
ardal.no
årdal.no
aremark.no
arendal.no
ås.no
aseral.no
åseral.no
asker.no
askim.no
askvoll.no
askoy.no
askøy.no
asnes.no
åsnes.no
audnedaln.no
aukra.no
aure.no
aurland.no
aurskog-holand.no
aurskog-høland.no
austevoll.no
austrheim.no
averoy.no
averøy.no
balestrand.no
ballangen.no
balat.no
bálát.no
balsfjord.no
bahccavuotna.no
báhccavuotna.no
bamble.no
bardu.no
beardu.no
beiarn.no
bajddar.no
bájddar.no
baidar.no
báidár.no
berg.no
bergen.no
berlevag.no
berlevåg.no
bearalvahki.no
bearalváhki.no
bindal.no
birkenes.no
bjarkoy.no
bjarkøy.no
bjerkreim.no
bjugn.no
bodo.no
bodø.no
badaddja.no
bådåddjå.no
budejju.no
bokn.no
bremanger.no
bronnoy.no
brønnøy.no
bygland.no
bykle.no
barum.no
bærum.no
bo.telemark.no
bø.telemark.no
bo.nordland.no
bø.nordland.no
bievat.no
bievát.no
bomlo.no
bømlo.no
batsfjord.no
båtsfjord.no
bahcavuotna.no
báhcavuotna.no
dovre.no
drammen.no
drangedal.no
dyroy.no
dyrøy.no
donna.no
dønna.no
eid.no
eidfjord.no
eidsberg.no
eidskog.no
eidsvoll.no
eigersund.no
elverum.no
enebakk.no
engerdal.no
etne.no
etnedal.no
evenes.no
evenassi.no
evenášši.no
evje-og-hornnes.no
farsund.no
fauske.no
fuossko.no
fuoisku.no
fedje.no
fet.no
finnoy.no
finnøy.no
fitjar.no
fjaler.no
fjell.no
flakstad.no
flatanger.no
flekkefjord.no
flesberg.no
flora.no
fla.no
flå.no
folldal.no
forsand.no
fosnes.no
frei.no
frogn.no
froland.no
frosta.no
frana.no
fræna.no
froya.no
frøya.no
fusa.no
fyresdal.no
forde.no
førde.no
gamvik.no
gangaviika.no
gáŋgaviika.no
gaular.no
gausdal.no
gildeskal.no
gildeskål.no
giske.no
gjemnes.no
gjerdrum.no
gjerstad.no
gjesdal.no
gjovik.no
gjøvik.no
gloppen.no
gol.no
gran.no
grane.no
granvin.no
gratangen.no
grimstad.no
grong.no
kraanghke.no
kråanghke.no
grue.no
gulen.no
hadsel.no
halden.no
halsa.no
hamar.no
hamaroy.no
habmer.no
hábmer.no
hapmir.no
hápmir.no
hammerfest.no
hammarfeasta.no
hámmárfeasta.no
haram.no
hareid.no
harstad.no
hasvik.no
aknoluokta.no
ákŋoluokta.no
hattfjelldal.no
aarborte.no
haugesund.no
hemne.no
hemnes.no
hemsedal.no
heroy.more-og-romsdal.no
herøy.møre-og-romsdal.no
heroy.nordland.no
herøy.nordland.no
hitra.no
hjartdal.no
hjelmeland.no
hobol.no
hobøl.no
hof.no
hol.no
hole.no
holmestrand.no
holtalen.no
holtålen.no
hornindal.no
horten.no
hurdal.no
hurum.no
hvaler.no
hyllestad.no
hagebostad.no
hægebostad.no
hoyanger.no
høyanger.no
hoylandet.no
høylandet.no
ha.no
hå.no
ibestad.no
inderoy.no
inderøy.no
iveland.no
jevnaker.no
jondal.no
jolster.no
jølster.no
karasjok.no
karasjohka.no
kárášjohka.no
karlsoy.no
galsa.no
gálsá.no
karmoy.no
karmøy.no
kautokeino.no
guovdageaidnu.no
klepp.no
klabu.no
klæbu.no
kongsberg.no
kongsvinger.no
kragero.no
kragerø.no
kristiansand.no
kristiansund.no
krodsherad.no
krødsherad.no
kvalsund.no
rahkkeravju.no
ráhkkerávju.no
kvam.no
kvinesdal.no
kvinnherad.no
kviteseid.no
kvitsoy.no
kvitsøy.no
kvafjord.no
kvæfjord.no
giehtavuoatna.no
kvanangen.no
kvænangen.no
navuotna.no
návuotna.no
kafjord.no
kåfjord.no
gaivuotna.no
gáivuotna.no
larvik.no
lavangen.no
lavagis.no
loabat.no
loabát.no
lebesby.no
davvesiida.no
leikanger.no
leirfjord.no
leka.no
leksvik.no
lenvik.no
leangaviika.no
leaŋgaviika.no
lesja.no
levanger.no
lier.no
lierne.no
lillehammer.no
lillesand.no
lindesnes.no
lindas.no
lindås.no
lom.no
loppa.no
lahppi.no
láhppi.no
lund.no
lunner.no
luroy.no
lurøy.no
luster.no
lyngdal.no
lyngen.no
ivgu.no
lardal.no
lerdal.no
lærdal.no
lodingen.no
lødingen.no
lorenskog.no
lørenskog.no
loten.no
løten.no
malvik.no
masoy.no
måsøy.no
muosat.no
muosát.no
mandal.no
marker.no
marnardal.no
masfjorden.no
meland.no
meldal.no
melhus.no
meloy.no
meløy.no
meraker.no
meråker.no
moareke.no
moåreke.no
midsund.no
midtre-gauldal.no
modalen.no
modum.no
molde.no
moskenes.no
moss.no
mosvik.no
malselv.no
målselv.no
malatvuopmi.no
málatvuopmi.no
namdalseid.no
aejrie.no
namsos.no
namsskogan.no
naamesjevuemie.no
nååmesjevuemie.no
laakesvuemie.no
nannestad.no
narvik.no
narviika.no
naustdal.no
nedre-eiker.no
nes.akershus.no
nes.buskerud.no
nesna.no
nesodden.no
nesseby.no
unjarga.no
unjárga.no
nesset.no
nissedal.no
nittedal.no
nord-aurdal.no
nord-fron.no
nord-odal.no
norddal.no
nordkapp.no
davvenjarga.no
davvenjárga.no
nordre-land.no
nordreisa.no
raisa.no
ráisa.no
nore-og-uvdal.no
notodden.no
naroy.no
nærøy.no
notteroy.no
nøtterøy.no
odda.no
oksnes.no
øksnes.no
oppdal.no
oppegard.no
oppegård.no
orkdal.no
orland.no
ørland.no
orskog.no
ørskog.no
orsta.no
ørsta.no
os.hedmark.no
os.hordaland.no
osen.no
osteroy.no
osterøy.no
ostre-toten.no
østre-toten.no
overhalla.no
ovre-eiker.no
øvre-eiker.no
oyer.no
øyer.no
oygarden.no
øygarden.no
oystre-slidre.no
øystre-slidre.no
porsanger.no
porsangu.no
porsáŋgu.no
porsgrunn.no
radoy.no
radøy.no
rakkestad.no
rana.no
ruovat.no
randaberg.no
rauma.no
rendalen.no
rennebu.no
rennesoy.no
rennesøy.no
rindal.no
ringebu.no
ringerike.no
ringsaker.no
rissa.no
risor.no
risør.no
roan.no
rollag.no
rygge.no
ralingen.no
rælingen.no
rodoy.no
rødøy.no
romskog.no
rømskog.no
roros.no
røros.no
rost.no
røst.no
royken.no
røyken.no
royrvik.no
røyrvik.no
rade.no
råde.no
salangen.no
siellak.no
saltdal.no
salat.no
sálát.no
sálat.no
samnanger.no
sande.more-og-romsdal.no
sande.møre-og-romsdal.no
sande.vestfold.no
sandefjord.no
sandnes.no
sandoy.no
sandøy.no
sarpsborg.no
sauda.no
sauherad.no
sel.no
selbu.no
selje.no
seljord.no
sigdal.no
siljan.no
sirdal.no
skaun.no
skedsmo.no
ski.no
skien.no
skiptvet.no
skjervoy.no
skjervøy.no
skierva.no
skiervá.no
skjak.no
skjåk.no
skodje.no
skanland.no
skånland.no
skanit.no
skánit.no
smola.no
smøla.no
snillfjord.no
snasa.no
snåsa.no
snoasa.no
snaase.no
snåase.no
sogndal.no
sokndal.no
sola.no
solund.no
songdalen.no
sortland.no
spydeberg.no
stange.no
stavanger.no
steigen.no
steinkjer.no
stjordal.no
stjørdal.no
stokke.no
stor-elvdal.no
stord.no
stordal.no
storfjord.no
omasvuotna.no
strand.no
stranda.no
stryn.no
sula.no
suldal.no
sund.no
sunndal.no
surnadal.no
sveio.no
svelvik.no
sykkylven.no
sogne.no
søgne.no
somna.no
sømna.no
sondre-land.no
søndre-land.no
sor-aurdal.no
sør-aurdal.no
sor-fron.no
sør-fron.no
sor-odal.no
sør-odal.no
sor-varanger.no
sør-varanger.no
matta-varjjat.no
mátta-várjjat.no
sorfold.no
sørfold.no
sorreisa.no
sørreisa.no
sorum.no
sørum.no
tana.no
deatnu.no
time.no
tingvoll.no
tinn.no
tjeldsund.no
dielddanuorri.no
tjome.no
tjøme.no
tokke.no
tolga.no
torsken.no
tranoy.no
tranøy.no
tromso.no
tromsø.no
tromsa.no
romsa.no
trondheim.no
troandin.no
trysil.no
trana.no
træna.no
trogstad.no
trøgstad.no
tvedestrand.no
tydal.no
tynset.no
tysfjord.no
divtasvuodna.no
divttasvuotna.no
tysnes.no
tysvar.no
tysvær.no
tonsberg.no
tønsberg.no
ullensaker.no
ullensvang.no
ulvik.no
utsira.no
vadso.no
vadsø.no
cahcesuolo.no
čáhcesuolo.no
vaksdal.no
valle.no
vang.no
vanylven.no
vardo.no
vardø.no
varggat.no
várggát.no
vefsn.no
vaapste.no
vega.no
vegarshei.no
vegårshei.no
vennesla.no
verdal.no
verran.no
vestby.no
vestnes.no
vestre-slidre.no
vestre-toten.no
vestvagoy.no
vestvågøy.no
vevelstad.no
vik.no
vikna.no
vindafjord.no
volda.no
voss.no
varoy.no
værøy.no
vagan.no
vågan.no
voagat.no
vagsoy.no
vågsøy.no
vaga.no
vågå.no
valer.ostfold.no
våler.østfold.no
valer.hedmark.no
våler.hedmark.no

// np : http://www.mos.com.np/register.html
*.np

// nr : http://cenpac.net.nr/dns/index.html
// Submitted by registry <technician@cenpac.net.nr>
nr
biz.nr
info.nr
gov.nr
edu.nr
org.nr
net.nr
com.nr

// nu : https://en.wikipedia.org/wiki/.nu
nu

// nz : https://en.wikipedia.org/wiki/.nz
// Submitted by registry <jay@nzrs.net.nz>
nz
ac.nz
co.nz
cri.nz
geek.nz
gen.nz
govt.nz
health.nz
iwi.nz
kiwi.nz
maori.nz
mil.nz
māori.nz
net.nz
org.nz
parliament.nz
school.nz

// om : https://en.wikipedia.org/wiki/.om
om
co.om
com.om
edu.om
gov.om
med.om
museum.om
net.om
org.om
pro.om

// onion : https://tools.ietf.org/html/rfc7686
onion

// org : https://en.wikipedia.org/wiki/.org
org

// pa : http://www.nic.pa/
// Some additional second level "domains" resolve directly as hostnames, such as
// pannet.pa, so we add a rule for "pa".
pa
ac.pa
gob.pa
com.pa
org.pa
sld.pa
edu.pa
net.pa
ing.pa
abo.pa
med.pa
nom.pa

// pe : https://www.nic.pe/InformeFinalComision.pdf
pe
edu.pe
gob.pe
nom.pe
mil.pe
org.pe
com.pe
net.pe

// pf : http://www.gobin.info/domainname/formulaire-pf.pdf
pf
com.pf
org.pf
edu.pf

// pg : https://en.wikipedia.org/wiki/.pg
*.pg

// ph : http://www.domains.ph/FAQ2.asp
// Submitted by registry <jed@email.com.ph>
ph
com.ph
net.ph
org.ph
gov.ph
edu.ph
ngo.ph
mil.ph
i.ph

// pk : http://pk5.pknic.net.pk/pk5/msgNamepk.PK
pk
com.pk
net.pk
edu.pk
org.pk
fam.pk
biz.pk
web.pk
gov.pk
gob.pk
gok.pk
gon.pk
gop.pk
gos.pk
info.pk

// pl http://www.dns.pl/english/index.html
// Submitted by registry
pl
com.pl
net.pl
org.pl
// pl functional domains (http://www.dns.pl/english/index.html)
aid.pl
agro.pl
atm.pl
auto.pl
biz.pl
edu.pl
gmina.pl
gsm.pl
info.pl
mail.pl
miasta.pl
media.pl
mil.pl
nieruchomosci.pl
nom.pl
pc.pl
powiat.pl
priv.pl
realestate.pl
rel.pl
sex.pl
shop.pl
sklep.pl
sos.pl
szkola.pl
targi.pl
tm.pl
tourism.pl
travel.pl
turystyka.pl
// Government domains
gov.pl
ap.gov.pl
ic.gov.pl
is.gov.pl
us.gov.pl
kmpsp.gov.pl
kppsp.gov.pl
kwpsp.gov.pl
psp.gov.pl
wskr.gov.pl
kwp.gov.pl
mw.gov.pl
ug.gov.pl
um.gov.pl
umig.gov.pl
ugim.gov.pl
upow.gov.pl
uw.gov.pl
starostwo.gov.pl
pa.gov.pl
po.gov.pl
psse.gov.pl
pup.gov.pl
rzgw.gov.pl
sa.gov.pl
so.gov.pl
sr.gov.pl
wsa.gov.pl
sko.gov.pl
uzs.gov.pl
wiih.gov.pl
winb.gov.pl
pinb.gov.pl
wios.gov.pl
witd.gov.pl
wzmiuw.gov.pl
piw.gov.pl
wiw.gov.pl
griw.gov.pl
wif.gov.pl
oum.gov.pl
sdn.gov.pl
zp.gov.pl
uppo.gov.pl
mup.gov.pl
wuoz.gov.pl
konsulat.gov.pl
oirm.gov.pl
// pl regional domains (http://www.dns.pl/english/index.html)
augustow.pl
babia-gora.pl
bedzin.pl
beskidy.pl
bialowieza.pl
bialystok.pl
bielawa.pl
bieszczady.pl
boleslawiec.pl
bydgoszcz.pl
bytom.pl
cieszyn.pl
czeladz.pl
czest.pl
dlugoleka.pl
elblag.pl
elk.pl
glogow.pl
gniezno.pl
gorlice.pl
grajewo.pl
ilawa.pl
jaworzno.pl
jelenia-gora.pl
jgora.pl
kalisz.pl
kazimierz-dolny.pl
karpacz.pl
kartuzy.pl
kaszuby.pl
katowice.pl
kepno.pl
ketrzyn.pl
klodzko.pl
kobierzyce.pl
kolobrzeg.pl
konin.pl
konskowola.pl
kutno.pl
lapy.pl
lebork.pl
legnica.pl
lezajsk.pl
limanowa.pl
lomza.pl
lowicz.pl
lubin.pl
lukow.pl
malbork.pl
malopolska.pl
mazowsze.pl
mazury.pl
mielec.pl
mielno.pl
mragowo.pl
naklo.pl
nowaruda.pl
nysa.pl
olawa.pl
olecko.pl
olkusz.pl
olsztyn.pl
opoczno.pl
opole.pl
ostroda.pl
ostroleka.pl
ostrowiec.pl
ostrowwlkp.pl
pila.pl
pisz.pl
podhale.pl
podlasie.pl
polkowice.pl
pomorze.pl
pomorskie.pl
prochowice.pl
pruszkow.pl
przeworsk.pl
pulawy.pl
radom.pl
rawa-maz.pl
rybnik.pl
rzeszow.pl
sanok.pl
sejny.pl
slask.pl
slupsk.pl
sosnowiec.pl
stalowa-wola.pl
skoczow.pl
starachowice.pl
stargard.pl
suwalki.pl
swidnica.pl
swiebodzin.pl
swinoujscie.pl
szczecin.pl
szczytno.pl
tarnobrzeg.pl
tgory.pl
turek.pl
tychy.pl
ustka.pl
walbrzych.pl
warmia.pl
warszawa.pl
waw.pl
wegrow.pl
wielun.pl
wlocl.pl
wloclawek.pl
wodzislaw.pl
wolomin.pl
wroclaw.pl
zachpomor.pl
zagan.pl
zarow.pl
zgora.pl
zgorzelec.pl

// pm : http://www.afnic.fr/medias/documents/AFNIC-naming-policy2012.pdf
pm

// pn : http://www.government.pn/PnRegistry/policies.htm
pn
gov.pn
co.pn
org.pn
edu.pn
net.pn

// post : https://en.wikipedia.org/wiki/.post
post

// pr : http://www.nic.pr/index.asp?f=1
pr
com.pr
net.pr
org.pr
gov.pr
edu.pr
isla.pr
pro.pr
biz.pr
info.pr
name.pr
// these aren't mentioned on nic.pr, but on https://en.wikipedia.org/wiki/.pr
est.pr
prof.pr
ac.pr

// pro : http://registry.pro/get-pro
pro
aaa.pro
aca.pro
acct.pro
avocat.pro
bar.pro
cpa.pro
eng.pro
jur.pro
law.pro
med.pro
recht.pro

// ps : https://en.wikipedia.org/wiki/.ps
// http://www.nic.ps/registration/policy.html#reg
ps
edu.ps
gov.ps
sec.ps
plo.ps
com.ps
org.ps
net.ps

// pt : http://online.dns.pt/dns/start_dns
pt
net.pt
gov.pt
org.pt
edu.pt
int.pt
publ.pt
com.pt
nome.pt

// pw : https://en.wikipedia.org/wiki/.pw
pw
co.pw
ne.pw
or.pw
ed.pw
go.pw
belau.pw

// py : http://www.nic.py/pautas.html#seccion_9
// Submitted by registry
py
com.py
coop.py
edu.py
gov.py
mil.py
net.py
org.py

// qa : http://domains.qa/en/
qa
com.qa
edu.qa
gov.qa
mil.qa
name.qa
net.qa
org.qa
sch.qa

// re : http://www.afnic.re/obtenir/chartes/nommage-re/annexe-descriptifs
re
asso.re
com.re
nom.re

// ro : http://www.rotld.ro/
ro
arts.ro
com.ro
firm.ro
info.ro
nom.ro
nt.ro
org.ro
rec.ro
store.ro
tm.ro
www.ro

// rs : https://www.rnids.rs/en/domains/national-domains
rs
ac.rs
co.rs
edu.rs
gov.rs
in.rs
org.rs

// ru : https://cctld.ru/en/domains/domens_ru/reserved/
ru
ac.ru
edu.ru
gov.ru
int.ru
mil.ru
test.ru

// rw : http://www.nic.rw/cgi-bin/policy.pl
rw
gov.rw
net.rw
edu.rw
ac.rw
com.rw
co.rw
int.rw
mil.rw
gouv.rw

// sa : http://www.nic.net.sa/
sa
com.sa
net.sa
org.sa
gov.sa
med.sa
pub.sa
edu.sa
sch.sa

// sb : http://www.sbnic.net.sb/
// Submitted by registry <lee.humphries@telekom.com.sb>
sb
com.sb
edu.sb
gov.sb
net.sb
org.sb

// sc : http://www.nic.sc/
sc
com.sc
gov.sc
net.sc
org.sc
edu.sc

// sd : http://www.isoc.sd/sudanic.isoc.sd/billing_pricing.htm
// Submitted by registry <admin@isoc.sd>
sd
com.sd
net.sd
org.sd
edu.sd
med.sd
tv.sd
gov.sd
info.sd

// se : https://en.wikipedia.org/wiki/.se
// Submitted by registry <patrik.wallstrom@iis.se>
se
a.se
ac.se
b.se
bd.se
brand.se
c.se
d.se
e.se
f.se
fh.se
fhsk.se
fhv.se
g.se
h.se
i.se
k.se
komforb.se
kommunalforbund.se
komvux.se
l.se
lanbib.se
m.se
n.se
naturbruksgymn.se
o.se
org.se
p.se
parti.se
pp.se
press.se
r.se
s.se
t.se
tm.se
u.se
w.se
x.se
y.se
z.se

// sg : http://www.nic.net.sg/page/registration-policies-procedures-and-guidelines
sg
com.sg
net.sg
org.sg
gov.sg
edu.sg
per.sg

// sh : http://www.nic.sh/registrar.html
sh
com.sh
net.sh
gov.sh
org.sh
mil.sh

// si : https://en.wikipedia.org/wiki/.si
si

// sj : No registrations at this time.
// Submitted by registry <jarle@uninett.no>
sj

// sk : https://en.wikipedia.org/wiki/.sk
// list of 2nd level domains ?
sk

// sl : http://www.nic.sl
// Submitted by registry <adam@neoip.com>
sl
com.sl
net.sl
edu.sl
gov.sl
org.sl

// sm : https://en.wikipedia.org/wiki/.sm
sm

// sn : https://en.wikipedia.org/wiki/.sn
sn
art.sn
com.sn
edu.sn
gouv.sn
org.sn
perso.sn
univ.sn

// so : http://www.soregistry.com/
so
com.so
net.so
org.so

// sr : https://en.wikipedia.org/wiki/.sr
sr

// st : http://www.nic.st/html/policyrules/
st
co.st
com.st
consulado.st
edu.st
embaixada.st
gov.st
mil.st
net.st
org.st
principe.st
saotome.st
store.st

// su : https://en.wikipedia.org/wiki/.su
su

// sv : http://www.svnet.org.sv/niveldos.pdf
sv
com.sv
edu.sv
gob.sv
org.sv
red.sv

// sx : https://en.wikipedia.org/wiki/.sx
// Submitted by registry <jcvignes@openregistry.com>
sx
gov.sx

// sy : https://en.wikipedia.org/wiki/.sy
// see also: http://www.gobin.info/domainname/sy.doc
sy
edu.sy
gov.sy
net.sy
mil.sy
com.sy
org.sy

// sz : https://en.wikipedia.org/wiki/.sz
// http://www.sispa.org.sz/
sz
co.sz
ac.sz
org.sz

// tc : https://en.wikipedia.org/wiki/.tc
tc

// td : https://en.wikipedia.org/wiki/.td
td

// tel: https://en.wikipedia.org/wiki/.tel
// http://www.telnic.org/
tel

// tf : https://en.wikipedia.org/wiki/.tf
tf

// tg : https://en.wikipedia.org/wiki/.tg
// http://www.nic.tg/
tg

// th : https://en.wikipedia.org/wiki/.th
// Submitted by registry <krit@thains.co.th>
th
ac.th
co.th
go.th
in.th
mi.th
net.th
or.th

// tj : http://www.nic.tj/policy.html
tj
ac.tj
biz.tj
co.tj
com.tj
edu.tj
go.tj
gov.tj
int.tj
mil.tj
name.tj
net.tj
nic.tj
org.tj
test.tj
web.tj

// tk : https://en.wikipedia.org/wiki/.tk
tk

// tl : https://en.wikipedia.org/wiki/.tl
tl
gov.tl

// tm : http://www.nic.tm/local.html
tm
com.tm
co.tm
org.tm
net.tm
nom.tm
gov.tm
mil.tm
edu.tm

// tn : https://en.wikipedia.org/wiki/.tn
// http://whois.ati.tn/
tn
com.tn
ens.tn
fin.tn
gov.tn
ind.tn
intl.tn
nat.tn
net.tn
org.tn
info.tn
perso.tn
tourism.tn
edunet.tn
rnrt.tn
rns.tn
rnu.tn
mincom.tn
agrinet.tn
defense.tn
turen.tn

// to : https://en.wikipedia.org/wiki/.to
// Submitted by registry <egullich@colo.to>
to
com.to
gov.to
net.to
org.to
edu.to
mil.to

// subTLDs: https://www.nic.tr/forms/eng/policies.pdf
//     and: https://www.nic.tr/forms/politikalar.pdf
// Submitted by <mehmetgurevin@gmail.com>
tr
com.tr
info.tr
biz.tr
net.tr
org.tr
web.tr
gen.tr
tv.tr
av.tr
dr.tr
bbs.tr
name.tr
tel.tr
gov.tr
bel.tr
pol.tr
mil.tr
k12.tr
edu.tr
kep.tr

// Used by Northern Cyprus
nc.tr

// Used by government agencies of Northern Cyprus
gov.nc.tr

// tt : http://www.nic.tt/
tt
co.tt
com.tt
org.tt
net.tt
biz.tt
info.tt
pro.tt
int.tt
coop.tt
jobs.tt
mobi.tt
travel.tt
museum.tt
aero.tt
name.tt
gov.tt
edu.tt

// tv : https://en.wikipedia.org/wiki/.tv
// Not listing any 2LDs as reserved since none seem to exist in practice,
// Wikipedia notwithstanding.
tv

// tw : https://en.wikipedia.org/wiki/.tw
tw
edu.tw
gov.tw
mil.tw
com.tw
net.tw
org.tw
idv.tw
game.tw
ebiz.tw
club.tw
網路.tw
組織.tw
商業.tw

// tz : http://www.tznic.or.tz/index.php/domains
// Submitted by registry <manager@tznic.or.tz>
tz
ac.tz
co.tz
go.tz
hotel.tz
info.tz
me.tz
mil.tz
mobi.tz
ne.tz
or.tz
sc.tz
tv.tz

// ua : https://hostmaster.ua/policy/?ua
// Submitted by registry <dk@cctld.ua>
ua
// ua 2LD
com.ua
edu.ua
gov.ua
in.ua
net.ua
org.ua
// ua geographic names
// https://hostmaster.ua/2ld/
cherkassy.ua
cherkasy.ua
chernigov.ua
chernihiv.ua
chernivtsi.ua
chernovtsy.ua
ck.ua
cn.ua
cr.ua
crimea.ua
cv.ua
dn.ua
dnepropetrovsk.ua
dnipropetrovsk.ua
dominic.ua
donetsk.ua
dp.ua
if.ua
ivano-frankivsk.ua
kh.ua
kharkiv.ua
kharkov.ua
kherson.ua
khmelnitskiy.ua
khmelnytskyi.ua
kiev.ua
kirovograd.ua
km.ua
kr.ua
krym.ua
ks.ua
kv.ua
kyiv.ua
lg.ua
lt.ua
lugansk.ua
lutsk.ua
lv.ua
lviv.ua
mk.ua
mykolaiv.ua
nikolaev.ua
od.ua
odesa.ua
odessa.ua
pl.ua
poltava.ua
rivne.ua
rovno.ua
rv.ua
sb.ua
sebastopol.ua
sevastopol.ua
sm.ua
sumy.ua
te.ua
ternopil.ua
uz.ua
uzhgorod.ua
vinnica.ua
vinnytsia.ua
vn.ua
volyn.ua
yalta.ua
zaporizhzhe.ua
zaporizhzhia.ua
zhitomir.ua
zhytomyr.ua
zp.ua
zt.ua

// ug : https://www.registry.co.ug/
ug
co.ug
or.ug
ac.ug
sc.ug
go.ug
ne.ug
com.ug
org.ug

// uk : https://en.wikipedia.org/wiki/.uk
// Submitted by registry <Michael.Daly@nominet.org.uk>
uk
ac.uk
co.uk
gov.uk
ltd.uk
me.uk
net.uk
nhs.uk
org.uk
plc.uk
police.uk
*.sch.uk

// us : https://en.wikipedia.org/wiki/.us
us
dni.us
fed.us
isa.us
kids.us
nsn.us
// us geographic names
ak.us
al.us
ar.us
as.us
az.us
ca.us
co.us
ct.us
dc.us
de.us
fl.us
ga.us
gu.us
hi.us
ia.us
id.us
il.us
in.us
ks.us
ky.us
la.us
ma.us
md.us
me.us
mi.us
mn.us
mo.us
ms.us
mt.us
nc.us
nd.us
ne.us
nh.us
nj.us
nm.us
nv.us
ny.us
oh.us
ok.us
or.us
pa.us
pr.us
ri.us
sc.us
sd.us
tn.us
tx.us
ut.us
vi.us
vt.us
va.us
wa.us
wi.us
wv.us
wy.us
// The registrar notes several more specific domains available in each state,
// such as state.*.us, dst.*.us, etc., but resolution of these is somewhat
// haphazard; in some states these domains resolve as addresses, while in others
// only subdomains are available, or even nothing at all. We include the
// most common ones where it's clear that different sites are different
// entities.
k12.ak.us
k12.al.us
k12.ar.us
k12.as.us
k12.az.us
k12.ca.us
k12.co.us
k12.ct.us
k12.dc.us
k12.de.us
k12.fl.us
k12.ga.us
k12.gu.us
// k12.hi.us  Bug 614565 - Hawaii has a state-wide DOE login
k12.ia.us
k12.id.us
k12.il.us
k12.in.us
k12.ks.us
k12.ky.us
k12.la.us
k12.ma.us
k12.md.us
k12.me.us
k12.mi.us
k12.mn.us
k12.mo.us
k12.ms.us
k12.mt.us
k12.nc.us
// k12.nd.us  Bug 1028347 - Removed at request of Travis Rosso <trossow@nd.gov>
k12.ne.us
k12.nh.us
k12.nj.us
k12.nm.us
k12.nv.us
k12.ny.us
k12.oh.us
k12.ok.us
k12.or.us
k12.pa.us
k12.pr.us
k12.ri.us
k12.sc.us
// k12.sd.us  Bug 934131 - Removed at request of James Booze <James.Booze@k12.sd.us>
k12.tn.us
k12.tx.us
k12.ut.us
k12.vi.us
k12.vt.us
k12.va.us
k12.wa.us
k12.wi.us
// k12.wv.us  Bug 947705 - Removed at request of Verne Britton <verne@wvnet.edu>
k12.wy.us
cc.ak.us
cc.al.us
cc.ar.us
cc.as.us
cc.az.us
cc.ca.us
cc.co.us
cc.ct.us
cc.dc.us
cc.de.us
cc.fl.us
cc.ga.us
cc.gu.us
cc.hi.us
cc.ia.us
cc.id.us
cc.il.us
cc.in.us
cc.ks.us
cc.ky.us
cc.la.us
cc.ma.us
cc.md.us
cc.me.us
cc.mi.us
cc.mn.us
cc.mo.us
cc.ms.us
cc.mt.us
cc.nc.us
cc.nd.us
cc.ne.us
cc.nh.us
cc.nj.us
cc.nm.us
cc.nv.us
cc.ny.us
cc.oh.us
cc.ok.us
cc.or.us
cc.pa.us
cc.pr.us
cc.ri.us
cc.sc.us
cc.sd.us
cc.tn.us
cc.tx.us
cc.ut.us
cc.vi.us
cc.vt.us
cc.va.us
cc.wa.us
cc.wi.us
cc.wv.us
cc.wy.us
lib.ak.us
lib.al.us
lib.ar.us
lib.as.us
lib.az.us
lib.ca.us
lib.co.us
lib.ct.us
lib.dc.us
// lib.de.us  Issue #243 - Moved to Private section at request of Ed Moore <Ed.Moore@lib.de.us>
lib.fl.us
lib.ga.us
lib.gu.us
lib.hi.us
lib.ia.us
lib.id.us
lib.il.us
lib.in.us
lib.ks.us
lib.ky.us
lib.la.us
lib.ma.us
lib.md.us
lib.me.us
lib.mi.us
lib.mn.us
lib.mo.us
lib.ms.us
lib.mt.us
lib.nc.us
lib.nd.us
lib.ne.us
lib.nh.us
lib.nj.us
lib.nm.us
lib.nv.us
lib.ny.us
lib.oh.us
lib.ok.us
lib.or.us
lib.pa.us
lib.pr.us
lib.ri.us
lib.sc.us
lib.sd.us
lib.tn.us
lib.tx.us
lib.ut.us
lib.vi.us
lib.vt.us
lib.va.us
lib.wa.us
lib.wi.us
// lib.wv.us  Bug 941670 - Removed at request of Larry W Arnold <arnold@wvlc.lib.wv.us>
lib.wy.us
// k12.ma.us contains school districts in Massachusetts. The 4LDs are
//  managed independently except for private (PVT), charter (CHTR) and
//  parochial (PAROCH) schools.  Those are delegated directly to the
//  5LD operators.   <k12-ma-hostmaster _ at _ rsuc.gweep.net>
pvt.k12.ma.us
chtr.k12.ma.us
paroch.k12.ma.us
// Merit Network, Inc. maintains the registry for =~ /(k12|cc|lib).mi.us/ and the following
//    see also: http://domreg.merit.edu
//    see also: whois -h whois.domreg.merit.edu help
ann-arbor.mi.us
cog.mi.us
dst.mi.us
eaton.mi.us
gen.mi.us
mus.mi.us
tec.mi.us
washtenaw.mi.us

// uy : http://www.nic.org.uy/
uy
com.uy
edu.uy
gub.uy
mil.uy
net.uy
org.uy

// uz : http://www.reg.uz/
uz
co.uz
com.uz
net.uz
org.uz

// va : https://en.wikipedia.org/wiki/.va
va

// vc : https://en.wikipedia.org/wiki/.vc
// Submitted by registry <kshah@ca.afilias.info>
vc
com.vc
net.vc
org.vc
gov.vc
mil.vc
edu.vc

// ve : https://registro.nic.ve/
// Submitted by registry
ve
arts.ve
co.ve
com.ve
e12.ve
edu.ve
firm.ve
gob.ve
gov.ve
info.ve
int.ve
mil.ve
net.ve
org.ve
rec.ve
store.ve
tec.ve
web.ve

// vg : https://en.wikipedia.org/wiki/.vg
vg

// vi : http://www.nic.vi/newdomainform.htm
// http://www.nic.vi/Domain_Rules/body_domain_rules.html indicates some other
// TLDs are "reserved", such as edu.vi and gov.vi, but doesn't actually say they
// are available for registration (which they do not seem to be).
vi
co.vi
com.vi
k12.vi
net.vi
org.vi

// vn : https://www.dot.vn/vnnic/vnnic/domainregistration.jsp
vn
com.vn
net.vn
org.vn
edu.vn
gov.vn
int.vn
ac.vn
biz.vn
info.vn
name.vn
pro.vn
health.vn

// vu : https://en.wikipedia.org/wiki/.vu
// http://www.vunic.vu/
vu
com.vu
edu.vu
net.vu
org.vu

// wf : http://www.afnic.fr/medias/documents/AFNIC-naming-policy2012.pdf
wf

// ws : https://en.wikipedia.org/wiki/.ws
// http://samoanic.ws/index.dhtml
ws
com.ws
net.ws
org.ws
gov.ws
edu.ws

// yt : http://www.afnic.fr/medias/documents/AFNIC-naming-policy2012.pdf
yt

// IDN ccTLDs
// When submitting patches, please maintain a sort by ISO 3166 ccTLD, then
// U-label, and follow this format:
// // A-Label ("<Latin renderings>", <language name>[, variant info]) : <ISO 3166 ccTLD>
// // [sponsoring org]
// U-Label

// xn--mgbaam7a8h ("Emerat", Arabic) : AE
// http://nic.ae/english/arabicdomain/rules.jsp
امارات

// xn--y9a3aq ("hye", Armenian) : AM
// ISOC AM (operated by .am Registry)
հայ

// xn--54b7fta0cc ("Bangla", Bangla) : BD
বাংলা

// xn--90ae ("bg", Bulgarian) : BG
бг

// xn--90ais ("bel", Belarusian/Russian Cyrillic) : BY
// Operated by .by registry
бел

// xn--fiqs8s ("Zhongguo/China", Chinese, Simplified) : CN
// CNNIC
// http://cnnic.cn/html/Dir/2005/10/11/3218.htm
中国

// xn--fiqz9s ("Zhongguo/China", Chinese, Traditional) : CN
// CNNIC
// http://cnnic.cn/html/Dir/2005/10/11/3218.htm
中國

// xn--lgbbat1ad8j ("Algeria/Al Jazair", Arabic) : DZ
الجزائر

// xn--wgbh1c ("Egypt/Masr", Arabic) : EG
// http://www.dotmasr.eg/
مصر

// xn--e1a4c ("eu", Cyrillic) : EU
ею

// xn--node ("ge", Georgian Mkhedruli) : GE
გე

// xn--qxam ("el", Greek) : GR
// Hellenic Ministry of Infrastructure, Transport, and Networks
ελ

// xn--j6w193g ("Hong Kong", Chinese) : HK
// https://www.hkirc.hk
// Submitted by registry <hk.tech@hkirc.hk>
// https://www.hkirc.hk/content.jsp?id=30#!/34
香港
公司.香港
教育.香港
政府.香港
個人.香港
網絡.香港
組織.香港

// xn--2scrj9c ("Bharat", Kannada) : IN
// India
ಭಾರತ

// xn--3hcrj9c ("Bharat", Oriya) : IN
// India
ଭାରତ

// xn--45br5cyl ("Bharatam", Assamese) : IN
// India
ভাৰত

// xn--h2breg3eve ("Bharatam", Sanskrit) : IN
// India
भारतम्

// xn--h2brj9c8c ("Bharot", Santali) : IN
// India
भारोत

// xn--mgbgu82a ("Bharat", Sindhi) : IN
// India
ڀارت

// xn--rvc1e0am3e ("Bharatam", Malayalam) : IN
// India
ഭാരതം

// xn--h2brj9c ("Bharat", Devanagari) : IN
// India
भारत

// xn--mgbbh1a ("Bharat", Kashmiri) : IN
// India
بارت

// xn--mgbbh1a71e ("Bharat", Arabic) : IN
// India
بھارت

// xn--fpcrj9c3d ("Bharat", Telugu) : IN
// India
భారత్

// xn--gecrj9c ("Bharat", Gujarati) : IN
// India
ભારત

// xn--s9brj9c ("Bharat", Gurmukhi) : IN
// India
ਭਾਰਤ

// xn--45brj9c ("Bharat", Bengali) : IN
// India
ভারত

// xn--xkc2dl3a5ee0h ("India", Tamil) : IN
// India
இந்தியா

// xn--mgba3a4f16a ("Iran", Persian) : IR
ایران

// xn--mgba3a4fra ("Iran", Arabic) : IR
ايران

// xn--mgbtx2b ("Iraq", Arabic) : IQ
// Communications and Media Commission
عراق

// xn--mgbayh7gpa ("al-Ordon", Arabic) : JO
// National Information Technology Center (NITC)
// Royal Scientific Society, Al-Jubeiha
الاردن

// xn--3e0b707e ("Republic of Korea", Hangul) : KR
한국

// xn--80ao21a ("Kaz", Kazakh) : KZ
қаз

// xn--fzc2c9e2c ("Lanka", Sinhalese-Sinhala) : LK
// http://nic.lk
ලංකා

// xn--xkc2al3hye2a ("Ilangai", Tamil) : LK
// http://nic.lk
இலங்கை

// xn--mgbc0a9azcg ("Morocco/al-Maghrib", Arabic) : MA
المغرب

// xn--d1alf ("mkd", Macedonian) : MK
// MARnet
мкд

// xn--l1acc ("mon", Mongolian) : MN
мон

// xn--mix891f ("Macao", Chinese, Traditional) : MO
// MONIC / HNET Asia (Registry Operator for .mo)
澳門

// xn--mix082f ("Macao", Chinese, Simplified) : MO
澳门

// xn--mgbx4cd0ab ("Malaysia", Malay) : MY
مليسيا

// xn--mgb9awbf ("Oman", Arabic) : OM
عمان

// xn--mgbai9azgqp6j ("Pakistan", Urdu/Arabic) : PK
پاکستان

// xn--mgbai9a5eva00b ("Pakistan", Urdu/Arabic, variant) : PK
پاكستان

// xn--ygbi2ammx ("Falasteen", Arabic) : PS
// The Palestinian National Internet Naming Authority (PNINA)
// http://www.pnina.ps
فلسطين

// xn--90a3ac ("srb", Cyrillic) : RS
// https://www.rnids.rs/en/domains/national-domains
срб
пр.срб
орг.срб
обр.срб
од.срб
упр.срб
ак.срб

// xn--p1ai ("rf", Russian-Cyrillic) : RU
// http://www.cctld.ru/en/docs/rulesrf.php
рф

// xn--wgbl6a ("Qatar", Arabic) : QA
// http://www.ict.gov.qa/
قطر

// xn--mgberp4a5d4ar ("AlSaudiah", Arabic) : SA
// http://www.nic.net.sa/
السعودية

// xn--mgberp4a5d4a87g ("AlSaudiah", Arabic, variant)  : SA
السعودیة

// xn--mgbqly7c0a67fbc ("AlSaudiah", Arabic, variant) : SA
السعودیۃ

// xn--mgbqly7cvafr ("AlSaudiah", Arabic, variant) : SA
السعوديه

// xn--mgbpl2fh ("sudan", Arabic) : SD
// Operated by .sd registry
سودان

// xn--yfro4i67o Singapore ("Singapore", Chinese) : SG
新加坡

// xn--clchc0ea0b2g2a9gcd ("Singapore", Tamil) : SG
சிங்கப்பூர்

// xn--ogbpf8fl ("Syria", Arabic) : SY
سورية

// xn--mgbtf8fl ("Syria", Arabic, variant) : SY
سوريا

// xn--o3cw4h ("Thai", Thai) : TH
// http://www.thnic.co.th
ไทย
ศึกษา.ไทย
ธุรกิจ.ไทย
รัฐบาล.ไทย
ทหาร.ไทย
เน็ต.ไทย
องค์กร.ไทย

// xn--pgbs0dh ("Tunisia", Arabic) : TN
// http://nic.tn
تونس

// xn--kpry57d ("Taiwan", Chinese, Traditional) : TW
// http://www.twnic.net/english/dn/dn_07a.htm
台灣

// xn--kprw13d ("Taiwan", Chinese, Simplified) : TW
// http://www.twnic.net/english/dn/dn_07a.htm
台湾

// xn--nnx388a ("Taiwan", Chinese, variant) : TW
臺灣

// xn--j1amh ("ukr", Cyrillic) : UA
укр

// xn--mgb2ddes ("AlYemen", Arabic) : YE
اليمن

// xxx : http://icmregistry.com
xxx

// ye : http://www.y.net.ye/services/domain_name.htm
*.ye

// za : http://www.zadna.org.za/content/page/domain-information
ac.za
agric.za
alt.za
co.za
edu.za
gov.za
grondar.za
law.za
mil.za
net.za
ngo.za
nis.za
nom.za
org.za
school.za
tm.za
web.za

// zm : https://zicta.zm/
// Submitted by registry <info@zicta.zm>
zm
ac.zm
biz.zm
co.zm
com.zm
edu.zm
gov.zm
info.zm
mil.zm
net.zm
org.zm
sch.zm

// zw : https://www.potraz.gov.zw/
// Confirmed by registry <bmtengwa@potraz.gov.zw> 2017-01-25
zw
ac.zw
co.zw
gov.zw
mil.zw
org.zw


// newGTLDs
// List of new gTLDs imported from https://newgtlds.icann.org/newgtlds.csv on 2018-05-08T19:40:37Z
// This list is auto-generated, don't edit it manually.

// aaa : 2015-02-26 American Automobile Association, Inc.
aaa

// aarp : 2015-05-21 AARP
aarp

// abarth : 2015-07-30 Fiat Chrysler Automobiles N.V.
abarth

// abb : 2014-10-24 ABB Ltd
abb

// abbott : 2014-07-24 Abbott Laboratories, Inc.
abbott

// abbvie : 2015-07-30 AbbVie Inc.
abbvie

// abc : 2015-07-30 Disney Enterprises, Inc.
abc

// able : 2015-06-25 Able Inc.
able

// abogado : 2014-04-24 Minds + Machines Group Limited
abogado

// abudhabi : 2015-07-30 Abu Dhabi Systems and Information Centre
abudhabi

// academy : 2013-11-07 Binky Moon, LLC
academy

// accenture : 2014-08-15 Accenture plc
accenture

// accountant : 2014-11-20 dot Accountant Limited
accountant

// accountants : 2014-03-20 Binky Moon, LLC
accountants

// aco : 2015-01-08 ACO Severin Ahlmann GmbH & Co. KG
aco

// active : 2014-05-01 Active Network, LLC
active

// actor : 2013-12-12 United TLD Holdco Ltd.
actor

// adac : 2015-07-16 Allgemeiner Deutscher Automobil-Club e.V. (ADAC)
adac

// ads : 2014-12-04 Charleston Road Registry Inc.
ads

// adult : 2014-10-16 ICM Registry AD LLC
adult

// aeg : 2015-03-19 Aktiebolaget Electrolux
aeg

// aetna : 2015-05-21 Aetna Life Insurance Company
aetna

// afamilycompany : 2015-07-23 Johnson Shareholdings, Inc.
afamilycompany

// afl : 2014-10-02 Australian Football League
afl

// africa : 2014-03-24 ZA Central Registry NPC trading as Registry.Africa
africa

// agakhan : 2015-04-23 Fondation Aga Khan (Aga Khan Foundation)
agakhan

// agency : 2013-11-14 Binky Moon, LLC
agency

// aig : 2014-12-18 American International Group, Inc.
aig

// aigo : 2015-08-06 aigo Digital Technology Co,Ltd.
aigo

// airbus : 2015-07-30 Airbus S.A.S.
airbus

// airforce : 2014-03-06 United TLD Holdco Ltd.
airforce

// airtel : 2014-10-24 Bharti Airtel Limited
airtel

// akdn : 2015-04-23 Fondation Aga Khan (Aga Khan Foundation)
akdn

// alfaromeo : 2015-07-31 Fiat Chrysler Automobiles N.V.
alfaromeo

// alibaba : 2015-01-15 Alibaba Group Holding Limited
alibaba

// alipay : 2015-01-15 Alibaba Group Holding Limited
alipay

// allfinanz : 2014-07-03 Allfinanz Deutsche Vermögensberatung Aktiengesellschaft
allfinanz

// allstate : 2015-07-31 Allstate Fire and Casualty Insurance Company
allstate

// ally : 2015-06-18 Ally Financial Inc.
ally

// alsace : 2014-07-02 Region Grand Est
alsace

// alstom : 2015-07-30 ALSTOM
alstom

// americanexpress : 2015-07-31 American Express Travel Related Services Company, Inc.
americanexpress

// americanfamily : 2015-07-23 AmFam, Inc.
americanfamily

// amex : 2015-07-31 American Express Travel Related Services Company, Inc.
amex

// amfam : 2015-07-23 AmFam, Inc.
amfam

// amica : 2015-05-28 Amica Mutual Insurance Company
amica

// amsterdam : 2014-07-24 Gemeente Amsterdam
amsterdam

// analytics : 2014-12-18 Campus IP LLC
analytics

// android : 2014-08-07 Charleston Road Registry Inc.
android

// anquan : 2015-01-08 QIHOO 360 TECHNOLOGY CO. LTD.
anquan

// anz : 2015-07-31 Australia and New Zealand Banking Group Limited
anz

// aol : 2015-09-17 Oath Inc.
aol

// apartments : 2014-12-11 Binky Moon, LLC
apartments

// app : 2015-05-14 Charleston Road Registry Inc.
app

// apple : 2015-05-14 Apple Inc.
apple

// aquarelle : 2014-07-24 Aquarelle.com
aquarelle

// arab : 2015-11-12 League of Arab States
arab

// aramco : 2014-11-20 Aramco Services Company
aramco

// archi : 2014-02-06 Afilias plc
archi

// army : 2014-03-06 United TLD Holdco Ltd.
army

// art : 2016-03-24 UK Creative Ideas Limited
art

// arte : 2014-12-11 Association Relative à la Télévision Européenne G.E.I.E.
arte

// asda : 2015-07-31 Wal-Mart Stores, Inc.
asda

// associates : 2014-03-06 Binky Moon, LLC
associates

// athleta : 2015-07-30 The Gap, Inc.
athleta

// attorney : 2014-03-20 United TLD Holdco Ltd.
attorney

// auction : 2014-03-20 United TLD Holdco Ltd.
auction

// audi : 2015-05-21 AUDI Aktiengesellschaft
audi

// audible : 2015-06-25 Amazon Registry Services, Inc.
audible

// audio : 2014-03-20 Uniregistry, Corp.
audio

// auspost : 2015-08-13 Australian Postal Corporation
auspost

// author : 2014-12-18 Amazon Registry Services, Inc.
author

// auto : 2014-11-13 Cars Registry Limited
auto

// autos : 2014-01-09 DERAutos, LLC
autos

// avianca : 2015-01-08 Aerovias del Continente Americano S.A. Avianca
avianca

// aws : 2015-06-25 Amazon Registry Services, Inc.
aws

// axa : 2013-12-19 AXA SA
axa

// azure : 2014-12-18 Microsoft Corporation
azure

// baby : 2015-04-09 Johnson & Johnson Services, Inc.
baby

// baidu : 2015-01-08 Baidu, Inc.
baidu

// banamex : 2015-07-30 Citigroup Inc.
banamex

// bananarepublic : 2015-07-31 The Gap, Inc.
bananarepublic

// band : 2014-06-12 United TLD Holdco Ltd.
band

// bank : 2014-09-25 fTLD Registry Services LLC
bank

// bar : 2013-12-12 Punto 2012 Sociedad Anonima Promotora de Inversion de Capital Variable
bar

// barcelona : 2014-07-24 Municipi de Barcelona
barcelona

// barclaycard : 2014-11-20 Barclays Bank PLC
barclaycard

// barclays : 2014-11-20 Barclays Bank PLC
barclays

// barefoot : 2015-06-11 Gallo Vineyards, Inc.
barefoot

// bargains : 2013-11-14 Binky Moon, LLC
bargains

// baseball : 2015-10-29 MLB Advanced Media DH, LLC
baseball

// basketball : 2015-08-20 Fédération Internationale de Basketball (FIBA)
basketball

// bauhaus : 2014-04-17 Werkhaus GmbH
bauhaus

// bayern : 2014-01-23 Bayern Connect GmbH
bayern

// bbc : 2014-12-18 British Broadcasting Corporation
bbc

// bbt : 2015-07-23 BB&T Corporation
bbt

// bbva : 2014-10-02 BANCO BILBAO VIZCAYA ARGENTARIA, S.A.
bbva

// bcg : 2015-04-02 The Boston Consulting Group, Inc.
bcg

// bcn : 2014-07-24 Municipi de Barcelona
bcn

// beats : 2015-05-14 Beats Electronics, LLC
beats

// beauty : 2015-12-03 L'Oréal
beauty

// beer : 2014-01-09 Minds + Machines Group Limited
beer

// bentley : 2014-12-18 Bentley Motors Limited
bentley

// berlin : 2013-10-31 dotBERLIN GmbH & Co. KG
berlin

// best : 2013-12-19 BestTLD Pty Ltd
best

// bestbuy : 2015-07-31 BBY Solutions, Inc.
bestbuy

// bet : 2015-05-07 Afilias plc
bet

// bharti : 2014-01-09 Bharti Enterprises (Holding) Private Limited
bharti

// bible : 2014-06-19 American Bible Society
bible

// bid : 2013-12-19 dot Bid Limited
bid

// bike : 2013-08-27 Binky Moon, LLC
bike

// bing : 2014-12-18 Microsoft Corporation
bing

// bingo : 2014-12-04 Binky Moon, LLC
bingo

// bio : 2014-03-06 Afilias plc
bio

// black : 2014-01-16 Afilias plc
black

// blackfriday : 2014-01-16 Uniregistry, Corp.
blackfriday

// blanco : 2015-07-16 BLANCO GmbH + Co KG
blanco

// blockbuster : 2015-07-30 Dish DBS Corporation
blockbuster

// blog : 2015-05-14 Knock Knock WHOIS There, LLC
blog

// bloomberg : 2014-07-17 Bloomberg IP Holdings LLC
bloomberg

// blue : 2013-11-07 Afilias plc
blue

// bms : 2014-10-30 Bristol-Myers Squibb Company
bms

// bmw : 2014-01-09 Bayerische Motoren Werke Aktiengesellschaft
bmw

// bnl : 2014-07-24 Banca Nazionale del Lavoro
bnl

// bnpparibas : 2014-05-29 BNP Paribas
bnpparibas

// boats : 2014-12-04 DERBoats, LLC
boats

// boehringer : 2015-07-09 Boehringer Ingelheim International GmbH
boehringer

// bofa : 2015-07-31 Bank of America Corporation
bofa

// bom : 2014-10-16 Núcleo de Informação e Coordenação do Ponto BR - NIC.br
bom

// bond : 2014-06-05 Bond University Limited
bond

// boo : 2014-01-30 Charleston Road Registry Inc.
boo

// book : 2015-08-27 Amazon Registry Services, Inc.
book

// booking : 2015-07-16 Booking.com B.V.
booking

// bosch : 2015-06-18 Robert Bosch GMBH
bosch

// bostik : 2015-05-28 Bostik SA
bostik

// boston : 2015-12-10 Boston TLD Management, LLC
boston

// bot : 2014-12-18 Amazon Registry Services, Inc.
bot

// boutique : 2013-11-14 Binky Moon, LLC
boutique

// box : 2015-11-12 NS1 Limited
box

// bradesco : 2014-12-18 Banco Bradesco S.A.
bradesco

// bridgestone : 2014-12-18 Bridgestone Corporation
bridgestone

// broadway : 2014-12-22 Celebrate Broadway, Inc.
broadway

// broker : 2014-12-11 Dotbroker Registry Limited
broker

// brother : 2015-01-29 Brother Industries, Ltd.
brother

// brussels : 2014-02-06 DNS.be vzw
brussels

// budapest : 2013-11-21 Minds + Machines Group Limited
budapest

// bugatti : 2015-07-23 Bugatti International SA
bugatti

// build : 2013-11-07 Plan Bee LLC
build

// builders : 2013-11-07 Binky Moon, LLC
builders

// business : 2013-11-07 Binky Moon, LLC
business

// buy : 2014-12-18 Amazon Registry Services, Inc.
buy

// buzz : 2013-10-02 DOTSTRATEGY CO.
buzz

// bzh : 2014-02-27 Association www.bzh
bzh

// cab : 2013-10-24 Binky Moon, LLC
cab

// cafe : 2015-02-11 Binky Moon, LLC
cafe

// cal : 2014-07-24 Charleston Road Registry Inc.
cal

// call : 2014-12-18 Amazon Registry Services, Inc.
call

// calvinklein : 2015-07-30 PVH gTLD Holdings LLC
calvinklein

// cam : 2016-04-21 AC Webconnecting Holding B.V.
cam

// camera : 2013-08-27 Binky Moon, LLC
camera

// camp : 2013-11-07 Binky Moon, LLC
camp

// cancerresearch : 2014-05-15 Australian Cancer Research Foundation
cancerresearch

// canon : 2014-09-12 Canon Inc.
canon

// capetown : 2014-03-24 ZA Central Registry NPC trading as ZA Central Registry
capetown

// capital : 2014-03-06 Binky Moon, LLC
capital

// capitalone : 2015-08-06 Capital One Financial Corporation
capitalone

// car : 2015-01-22 Cars Registry Limited
car

// caravan : 2013-12-12 Caravan International, Inc.
caravan

// cards : 2013-12-05 Binky Moon, LLC
cards

// care : 2014-03-06 Binky Moon, LLC
care

// career : 2013-10-09 dotCareer LLC
career

// careers : 2013-10-02 Binky Moon, LLC
careers

// cars : 2014-11-13 Cars Registry Limited
cars

// cartier : 2014-06-23 Richemont DNS Inc.
cartier

// casa : 2013-11-21 Minds + Machines Group Limited
casa

// case : 2015-09-03 CNH Industrial N.V.
case

// caseih : 2015-09-03 CNH Industrial N.V.
caseih

// cash : 2014-03-06 Binky Moon, LLC
cash

// casino : 2014-12-18 Binky Moon, LLC
casino

// catering : 2013-12-05 Binky Moon, LLC
catering

// catholic : 2015-10-21 Pontificium Consilium de Comunicationibus Socialibus (PCCS) (Pontifical Council for Social Communication)
catholic

// cba : 2014-06-26 COMMONWEALTH BANK OF AUSTRALIA
cba

// cbn : 2014-08-22 The Christian Broadcasting Network, Inc.
cbn

// cbre : 2015-07-02 CBRE, Inc.
cbre

// cbs : 2015-08-06 CBS Domains Inc.
cbs

// ceb : 2015-04-09 The Corporate Executive Board Company
ceb

// center : 2013-11-07 Binky Moon, LLC
center

// ceo : 2013-11-07 CEOTLD Pty Ltd
ceo

// cern : 2014-06-05 European Organization for Nuclear Research ("CERN")
cern

// cfa : 2014-08-28 CFA Institute
cfa

// cfd : 2014-12-11 DotCFD Registry Limited
cfd

// chanel : 2015-04-09 Chanel International B.V.
chanel

// channel : 2014-05-08 Charleston Road Registry Inc.
channel

// charity : 2018-04-11 Corn Lake, LLC
charity

// chase : 2015-04-30 JPMorgan Chase Bank, National Association
chase

// chat : 2014-12-04 Binky Moon, LLC
chat

// cheap : 2013-11-14 Binky Moon, LLC
cheap

// chintai : 2015-06-11 CHINTAI Corporation
chintai

// christmas : 2013-11-21 Uniregistry, Corp.
christmas

// chrome : 2014-07-24 Charleston Road Registry Inc.
chrome

// chrysler : 2015-07-30 FCA US LLC.
chrysler

// church : 2014-02-06 Binky Moon, LLC
church

// cipriani : 2015-02-19 Hotel Cipriani Srl
cipriani

// circle : 2014-12-18 Amazon Registry Services, Inc.
circle

// cisco : 2014-12-22 Cisco Technology, Inc.
cisco

// citadel : 2015-07-23 Citadel Domain LLC
citadel

// citi : 2015-07-30 Citigroup Inc.
citi

// citic : 2014-01-09 CITIC Group Corporation
citic

// city : 2014-05-29 Binky Moon, LLC
city

// cityeats : 2014-12-11 Lifestyle Domain Holdings, Inc.
cityeats

// claims : 2014-03-20 Binky Moon, LLC
claims

// cleaning : 2013-12-05 Binky Moon, LLC
cleaning

// click : 2014-06-05 Uniregistry, Corp.
click

// clinic : 2014-03-20 Binky Moon, LLC
clinic

// clinique : 2015-10-01 The Estée Lauder Companies Inc.
clinique

// clothing : 2013-08-27 Binky Moon, LLC
clothing

// cloud : 2015-04-16 Aruba PEC S.p.A.
cloud

// club : 2013-11-08 .CLUB DOMAINS, LLC
club

// clubmed : 2015-06-25 Club Méditerranée S.A.
clubmed

// coach : 2014-10-09 Binky Moon, LLC
coach

// codes : 2013-10-31 Binky Moon, LLC
codes

// coffee : 2013-10-17 Binky Moon, LLC
coffee

// college : 2014-01-16 XYZ.COM LLC
college

// cologne : 2014-02-05 punkt.wien GmbH
cologne

// comcast : 2015-07-23 Comcast IP Holdings I, LLC
comcast

// commbank : 2014-06-26 COMMONWEALTH BANK OF AUSTRALIA
commbank

// community : 2013-12-05 Binky Moon, LLC
community

// company : 2013-11-07 Binky Moon, LLC
company

// compare : 2015-10-08 iSelect Ltd
compare

// computer : 2013-10-24 Binky Moon, LLC
computer

// comsec : 2015-01-08 VeriSign, Inc.
comsec

// condos : 2013-12-05 Binky Moon, LLC
condos

// construction : 2013-09-16 Binky Moon, LLC
construction

// consulting : 2013-12-05 United TLD Holdco Ltd.
consulting

// contact : 2015-01-08 Top Level Spectrum, Inc.
contact

// contractors : 2013-09-10 Binky Moon, LLC
contractors

// cooking : 2013-11-21 Minds + Machines Group Limited
cooking

// cookingchannel : 2015-07-02 Lifestyle Domain Holdings, Inc.
cookingchannel

// cool : 2013-11-14 Binky Moon, LLC
cool

// corsica : 2014-09-25 Collectivité de Corse
corsica

// country : 2013-12-19 DotCountry LLC
country

// coupon : 2015-02-26 Amazon Registry Services, Inc.
coupon

// coupons : 2015-03-26 Binky Moon, LLC
coupons

// courses : 2014-12-04 OPEN UNIVERSITIES AUSTRALIA PTY LTD
courses

// credit : 2014-03-20 Binky Moon, LLC
credit

// creditcard : 2014-03-20 Binky Moon, LLC
creditcard

// creditunion : 2015-01-22 CUNA Performance Resources, LLC
creditunion

// cricket : 2014-10-09 dot Cricket Limited
cricket

// crown : 2014-10-24 Crown Equipment Corporation
crown

// crs : 2014-04-03 Federated Co-operatives Limited
crs

// cruise : 2015-12-10 Viking River Cruises (Bermuda) Ltd.
cruise

// cruises : 2013-12-05 Binky Moon, LLC
cruises

// csc : 2014-09-25 Alliance-One Services, Inc.
csc

// cuisinella : 2014-04-03 SALM S.A.S.
cuisinella

// cymru : 2014-05-08 Nominet UK
cymru

// cyou : 2015-01-22 Beijing Gamease Age Digital Technology Co., Ltd.
cyou

// dabur : 2014-02-06 Dabur India Limited
dabur

// dad : 2014-01-23 Charleston Road Registry Inc.
dad

// dance : 2013-10-24 United TLD Holdco Ltd.
dance

// data : 2016-06-02 Dish DBS Corporation
data

// date : 2014-11-20 dot Date Limited
date

// dating : 2013-12-05 Binky Moon, LLC
dating

// datsun : 2014-03-27 NISSAN MOTOR CO., LTD.
datsun

// day : 2014-01-30 Charleston Road Registry Inc.
day

// dclk : 2014-11-20 Charleston Road Registry Inc.
dclk

// dds : 2015-05-07 Minds + Machines Group Limited
dds

// deal : 2015-06-25 Amazon Registry Services, Inc.
deal

// dealer : 2014-12-22 Dealer Dot Com, Inc.
dealer

// deals : 2014-05-22 Binky Moon, LLC
deals

// degree : 2014-03-06 United TLD Holdco Ltd.
degree

// delivery : 2014-09-11 Binky Moon, LLC
delivery

// dell : 2014-10-24 Dell Inc.
dell

// deloitte : 2015-07-31 Deloitte Touche Tohmatsu
deloitte

// delta : 2015-02-19 Delta Air Lines, Inc.
delta

// democrat : 2013-10-24 United TLD Holdco Ltd.
democrat

// dental : 2014-03-20 Binky Moon, LLC
dental

// dentist : 2014-03-20 United TLD Holdco Ltd.
dentist

// desi : 2013-11-14 Desi Networks LLC
desi

// design : 2014-11-07 Top Level Design, LLC
design

// dev : 2014-10-16 Charleston Road Registry Inc.
dev

// dhl : 2015-07-23 Deutsche Post AG
dhl

// diamonds : 2013-09-22 Binky Moon, LLC
diamonds

// diet : 2014-06-26 Uniregistry, Corp.
diet

// digital : 2014-03-06 Binky Moon, LLC
digital

// direct : 2014-04-10 Binky Moon, LLC
direct

// directory : 2013-09-20 Binky Moon, LLC
directory

// discount : 2014-03-06 Binky Moon, LLC
discount

// discover : 2015-07-23 Discover Financial Services
discover

// dish : 2015-07-30 Dish DBS Corporation
dish

// diy : 2015-11-05 Lifestyle Domain Holdings, Inc.
diy

// dnp : 2013-12-13 Dai Nippon Printing Co., Ltd.
dnp

// docs : 2014-10-16 Charleston Road Registry Inc.
docs

// doctor : 2016-06-02 Binky Moon, LLC
doctor

// dodge : 2015-07-30 FCA US LLC.
dodge

// dog : 2014-12-04 Binky Moon, LLC
dog

// doha : 2014-09-18 Communications Regulatory Authority (CRA)
doha

// domains : 2013-10-17 Binky Moon, LLC
domains

// dot : 2015-05-21 Dish DBS Corporation
dot

// download : 2014-11-20 dot Support Limited
download

// drive : 2015-03-05 Charleston Road Registry Inc.
drive

// dtv : 2015-06-04 Dish DBS Corporation
dtv

// dubai : 2015-01-01 Dubai Smart Government Department
dubai

// duck : 2015-07-23 Johnson Shareholdings, Inc.
duck

// dunlop : 2015-07-02 The Goodyear Tire & Rubber Company
dunlop

// duns : 2015-08-06 The Dun & Bradstreet Corporation
duns

// dupont : 2015-06-25 E. I. du Pont de Nemours and Company
dupont

// durban : 2014-03-24 ZA Central Registry NPC trading as ZA Central Registry
durban

// dvag : 2014-06-23 Deutsche Vermögensberatung Aktiengesellschaft DVAG
dvag

// dvr : 2016-05-26 Hughes Satellite Systems Corporation
dvr

// earth : 2014-12-04 Interlink Co., Ltd.
earth

// eat : 2014-01-23 Charleston Road Registry Inc.
eat

// eco : 2016-07-08 Big Room Inc.
eco

// edeka : 2014-12-18 EDEKA Verband kaufmännischer Genossenschaften e.V.
edeka

// education : 2013-11-07 Binky Moon, LLC
education

// email : 2013-10-31 Binky Moon, LLC
email

// emerck : 2014-04-03 Merck KGaA
emerck

// energy : 2014-09-11 Binky Moon, LLC
energy

// engineer : 2014-03-06 United TLD Holdco Ltd.
engineer

// engineering : 2014-03-06 Binky Moon, LLC
engineering

// enterprises : 2013-09-20 Binky Moon, LLC
enterprises

// epost : 2015-07-23 Deutsche Post AG
epost

// epson : 2014-12-04 Seiko Epson Corporation
epson

// equipment : 2013-08-27 Binky Moon, LLC
equipment

// ericsson : 2015-07-09 Telefonaktiebolaget L M Ericsson
ericsson

// erni : 2014-04-03 ERNI Group Holding AG
erni

// esq : 2014-05-08 Charleston Road Registry Inc.
esq

// estate : 2013-08-27 Binky Moon, LLC
estate

// esurance : 2015-07-23 Esurance Insurance Company
esurance

// etisalat : 2015-09-03 Emirates Telecommunications Corporation (trading as Etisalat)
etisalat

// eurovision : 2014-04-24 European Broadcasting Union (EBU)
eurovision

// eus : 2013-12-12 Puntueus Fundazioa
eus

// events : 2013-12-05 Binky Moon, LLC
events

// everbank : 2014-05-15 EverBank
everbank

// exchange : 2014-03-06 Binky Moon, LLC
exchange

// expert : 2013-11-21 Binky Moon, LLC
expert

// exposed : 2013-12-05 Binky Moon, LLC
exposed

// express : 2015-02-11 Binky Moon, LLC
express

// extraspace : 2015-05-14 Extra Space Storage LLC
extraspace

// fage : 2014-12-18 Fage International S.A.
fage

// fail : 2014-03-06 Binky Moon, LLC
fail

// fairwinds : 2014-11-13 FairWinds Partners, LLC
fairwinds

// faith : 2014-11-20 dot Faith Limited
faith

// family : 2015-04-02 United TLD Holdco Ltd.
family

// fan : 2014-03-06 Asiamix Digital Limited
fan

// fans : 2014-11-07 Asiamix Digital Limited
fans

// farm : 2013-11-07 Binky Moon, LLC
farm

// farmers : 2015-07-09 Farmers Insurance Exchange
farmers

// fashion : 2014-07-03 Minds + Machines Group Limited
fashion

// fast : 2014-12-18 Amazon Registry Services, Inc.
fast

// fedex : 2015-08-06 Federal Express Corporation
fedex

// feedback : 2013-12-19 Top Level Spectrum, Inc.
feedback

// ferrari : 2015-07-31 Fiat Chrysler Automobiles N.V.
ferrari

// ferrero : 2014-12-18 Ferrero Trading Lux S.A.
ferrero

// fiat : 2015-07-31 Fiat Chrysler Automobiles N.V.
fiat

// fidelity : 2015-07-30 Fidelity Brokerage Services LLC
fidelity

// fido : 2015-08-06 Rogers Communications Canada Inc.
fido

// film : 2015-01-08 Motion Picture Domain Registry Pty Ltd
film

// final : 2014-10-16 Núcleo de Informação e Coordenação do Ponto BR - NIC.br
final

// finance : 2014-03-20 Binky Moon, LLC
finance

// financial : 2014-03-06 Binky Moon, LLC
financial

// fire : 2015-06-25 Amazon Registry Services, Inc.
fire

// firestone : 2014-12-18 Bridgestone Licensing Services, Inc
firestone

// firmdale : 2014-03-27 Firmdale Holdings Limited
firmdale

// fish : 2013-12-12 Binky Moon, LLC
fish

// fishing : 2013-11-21 Minds + Machines Group Limited
fishing

// fit : 2014-11-07 Minds + Machines Group Limited
fit

// fitness : 2014-03-06 Binky Moon, LLC
fitness

// flickr : 2015-04-02 Yahoo! Domain Services Inc.
flickr

// flights : 2013-12-05 Binky Moon, LLC
flights

// flir : 2015-07-23 FLIR Systems, Inc.
flir

// florist : 2013-11-07 Binky Moon, LLC
florist

// flowers : 2014-10-09 Uniregistry, Corp.
flowers

// fly : 2014-05-08 Charleston Road Registry Inc.
fly

// foo : 2014-01-23 Charleston Road Registry Inc.
foo

// food : 2016-04-21 Lifestyle Domain Holdings, Inc.
food

// foodnetwork : 2015-07-02 Lifestyle Domain Holdings, Inc.
foodnetwork

// football : 2014-12-18 Binky Moon, LLC
football

// ford : 2014-11-13 Ford Motor Company
ford

// forex : 2014-12-11 Dotforex Registry Limited
forex

// forsale : 2014-05-22 United TLD Holdco Ltd.
forsale

// forum : 2015-04-02 Fegistry, LLC
forum

// foundation : 2013-12-05 Binky Moon, LLC
foundation

// fox : 2015-09-11 FOX Registry, LLC
fox

// free : 2015-12-10 Amazon Registry Services, Inc.
free

// fresenius : 2015-07-30 Fresenius Immobilien-Verwaltungs-GmbH
fresenius

// frl : 2014-05-15 FRLregistry B.V.
frl

// frogans : 2013-12-19 OP3FT
frogans

// frontdoor : 2015-07-02 Lifestyle Domain Holdings, Inc.
frontdoor

// frontier : 2015-02-05 Frontier Communications Corporation
frontier

// ftr : 2015-07-16 Frontier Communications Corporation
ftr

// fujitsu : 2015-07-30 Fujitsu Limited
fujitsu

// fujixerox : 2015-07-23 Xerox DNHC LLC
fujixerox

// fun : 2016-01-14 DotSpace Inc.
fun

// fund : 2014-03-20 Binky Moon, LLC
fund

// furniture : 2014-03-20 Binky Moon, LLC
furniture

// futbol : 2013-09-20 United TLD Holdco Ltd.
futbol

// fyi : 2015-04-02 Binky Moon, LLC
fyi

// gal : 2013-11-07 Asociación puntoGAL
gal

// gallery : 2013-09-13 Binky Moon, LLC
gallery

// gallo : 2015-06-11 Gallo Vineyards, Inc.
gallo

// gallup : 2015-02-19 Gallup, Inc.
gallup

// game : 2015-05-28 Uniregistry, Corp.
game

// games : 2015-05-28 United TLD Holdco Ltd.
games

// gap : 2015-07-31 The Gap, Inc.
gap

// garden : 2014-06-26 Minds + Machines Group Limited
garden

// gbiz : 2014-07-17 Charleston Road Registry Inc.
gbiz

// gdn : 2014-07-31 Joint Stock Company "Navigation-information systems"
gdn

// gea : 2014-12-04 GEA Group Aktiengesellschaft
gea

// gent : 2014-01-23 COMBELL NV
gent

// genting : 2015-03-12 Resorts World Inc Pte. Ltd.
genting

// george : 2015-07-31 Wal-Mart Stores, Inc.
george

// ggee : 2014-01-09 GMO Internet, Inc.
ggee

// gift : 2013-10-17 DotGift, LLC
gift

// gifts : 2014-07-03 Binky Moon, LLC
gifts

// gives : 2014-03-06 United TLD Holdco Ltd.
gives

// giving : 2014-11-13 Giving Limited
giving

// glade : 2015-07-23 Johnson Shareholdings, Inc.
glade

// glass : 2013-11-07 Binky Moon, LLC
glass

// gle : 2014-07-24 Charleston Road Registry Inc.
gle

// global : 2014-04-17 Dot Global Domain Registry Limited
global

// globo : 2013-12-19 Globo Comunicação e Participações S.A
globo

// gmail : 2014-05-01 Charleston Road Registry Inc.
gmail

// gmbh : 2016-01-29 Binky Moon, LLC
gmbh

// gmo : 2014-01-09 GMO Internet Pte. Ltd.
gmo

// gmx : 2014-04-24 1&1 Mail & Media GmbH
gmx

// godaddy : 2015-07-23 Go Daddy East, LLC
godaddy

// gold : 2015-01-22 Binky Moon, LLC
gold

// goldpoint : 2014-11-20 YODOBASHI CAMERA CO.,LTD.
goldpoint

// golf : 2014-12-18 Binky Moon, LLC
golf

// goo : 2014-12-18 NTT Resonant Inc.
goo

// goodyear : 2015-07-02 The Goodyear Tire & Rubber Company
goodyear

// goog : 2014-11-20 Charleston Road Registry Inc.
goog

// google : 2014-07-24 Charleston Road Registry Inc.
google

// gop : 2014-01-16 Republican State Leadership Committee, Inc.
gop

// got : 2014-12-18 Amazon Registry Services, Inc.
got

// grainger : 2015-05-07 Grainger Registry Services, LLC
grainger

// graphics : 2013-09-13 Binky Moon, LLC
graphics

// gratis : 2014-03-20 Binky Moon, LLC
gratis

// green : 2014-05-08 Afilias plc
green

// gripe : 2014-03-06 Binky Moon, LLC
gripe

// grocery : 2016-06-16 Wal-Mart Stores, Inc.
grocery

// group : 2014-08-15 Binky Moon, LLC
group

// guardian : 2015-07-30 The Guardian Life Insurance Company of America
guardian

// gucci : 2014-11-13 Guccio Gucci S.p.a.
gucci

// guge : 2014-08-28 Charleston Road Registry Inc.
guge

// guide : 2013-09-13 Binky Moon, LLC
guide

// guitars : 2013-11-14 Uniregistry, Corp.
guitars

// guru : 2013-08-27 Binky Moon, LLC
guru

// hair : 2015-12-03 L'Oréal
hair

// hamburg : 2014-02-20 Hamburg Top-Level-Domain GmbH
hamburg

// hangout : 2014-11-13 Charleston Road Registry Inc.
hangout

// haus : 2013-12-05 United TLD Holdco Ltd.
haus

// hbo : 2015-07-30 HBO Registry Services, Inc.
hbo

// hdfc : 2015-07-30 HOUSING DEVELOPMENT FINANCE CORPORATION LIMITED
hdfc

// hdfcbank : 2015-02-12 HDFC Bank Limited
hdfcbank

// health : 2015-02-11 DotHealth, LLC
health

// healthcare : 2014-06-12 Binky Moon, LLC
healthcare

// help : 2014-06-26 Uniregistry, Corp.
help

// helsinki : 2015-02-05 City of Helsinki
helsinki

// here : 2014-02-06 Charleston Road Registry Inc.
here

// hermes : 2014-07-10 HERMES INTERNATIONAL
hermes

// hgtv : 2015-07-02 Lifestyle Domain Holdings, Inc.
hgtv

// hiphop : 2014-03-06 Uniregistry, Corp.
hiphop

// hisamitsu : 2015-07-16 Hisamitsu Pharmaceutical Co.,Inc.
hisamitsu

// hitachi : 2014-10-31 Hitachi, Ltd.
hitachi

// hiv : 2014-03-13 Uniregistry, Corp.
hiv

// hkt : 2015-05-14 PCCW-HKT DataCom Services Limited
hkt

// hockey : 2015-03-19 Binky Moon, LLC
hockey

// holdings : 2013-08-27 Binky Moon, LLC
holdings

// holiday : 2013-11-07 Binky Moon, LLC
holiday

// homedepot : 2015-04-02 Home Depot Product Authority, LLC
homedepot

// homegoods : 2015-07-16 The TJX Companies, Inc.
homegoods

// homes : 2014-01-09 DERHomes, LLC
homes

// homesense : 2015-07-16 The TJX Companies, Inc.
homesense

// honda : 2014-12-18 Honda Motor Co., Ltd.
honda

// honeywell : 2015-07-23 Honeywell GTLD LLC
honeywell

// horse : 2013-11-21 Minds + Machines Group Limited
horse

// hospital : 2016-10-20 Binky Moon, LLC
hospital

// host : 2014-04-17 DotHost Inc.
host

// hosting : 2014-05-29 Uniregistry, Corp.
hosting

// hot : 2015-08-27 Amazon Registry Services, Inc.
hot

// hoteles : 2015-03-05 Travel Reservations SRL
hoteles

// hotels : 2016-04-07 Booking.com B.V.
hotels

// hotmail : 2014-12-18 Microsoft Corporation
hotmail

// house : 2013-11-07 Binky Moon, LLC
house

// how : 2014-01-23 Charleston Road Registry Inc.
how

// hsbc : 2014-10-24 HSBC Global Services (UK) Limited
hsbc

// hughes : 2015-07-30 Hughes Satellite Systems Corporation
hughes

// hyatt : 2015-07-30 Hyatt GTLD, L.L.C.
hyatt

// hyundai : 2015-07-09 Hyundai Motor Company
hyundai

// ibm : 2014-07-31 International Business Machines Corporation
ibm

// icbc : 2015-02-19 Industrial and Commercial Bank of China Limited
icbc

// ice : 2014-10-30 IntercontinentalExchange, Inc.
ice

// icu : 2015-01-08 ShortDot SA
icu

// ieee : 2015-07-23 IEEE Global LLC
ieee

// ifm : 2014-01-30 ifm electronic gmbh
ifm

// ikano : 2015-07-09 Ikano S.A.
ikano

// imamat : 2015-08-06 Fondation Aga Khan (Aga Khan Foundation)
imamat

// imdb : 2015-06-25 Amazon Registry Services, Inc.
imdb

// immo : 2014-07-10 Binky Moon, LLC
immo

// immobilien : 2013-11-07 United TLD Holdco Ltd.
immobilien

// inc : 2018-03-10 GTLD Limited
inc

// industries : 2013-12-05 Binky Moon, LLC
industries

// infiniti : 2014-03-27 NISSAN MOTOR CO., LTD.
infiniti

// ing : 2014-01-23 Charleston Road Registry Inc.
ing

// ink : 2013-12-05 Top Level Design, LLC
ink

// institute : 2013-11-07 Binky Moon, LLC
institute

// insurance : 2015-02-19 fTLD Registry Services LLC
insurance

// insure : 2014-03-20 Binky Moon, LLC
insure

// intel : 2015-08-06 Intel Corporation
intel

// international : 2013-11-07 Binky Moon, LLC
international

// intuit : 2015-07-30 Intuit Administrative Services, Inc.
intuit

// investments : 2014-03-20 Binky Moon, LLC
investments

// ipiranga : 2014-08-28 Ipiranga Produtos de Petroleo S.A.
ipiranga

// irish : 2014-08-07 Binky Moon, LLC
irish

// iselect : 2015-02-11 iSelect Ltd
iselect

// ismaili : 2015-08-06 Fondation Aga Khan (Aga Khan Foundation)
ismaili

// ist : 2014-08-28 Istanbul Metropolitan Municipality
ist

// istanbul : 2014-08-28 Istanbul Metropolitan Municipality
istanbul

// itau : 2014-10-02 Itau Unibanco Holding S.A.
itau

// itv : 2015-07-09 ITV Services Limited
itv

// iveco : 2015-09-03 CNH Industrial N.V.
iveco

// jaguar : 2014-11-13 Jaguar Land Rover Ltd
jaguar

// java : 2014-06-19 Oracle Corporation
java

// jcb : 2014-11-20 JCB Co., Ltd.
jcb

// jcp : 2015-04-23 JCP Media, Inc.
jcp

// jeep : 2015-07-30 FCA US LLC.
jeep

// jetzt : 2014-01-09 Binky Moon, LLC
jetzt

// jewelry : 2015-03-05 Binky Moon, LLC
jewelry

// jio : 2015-04-02 Reliance Industries Limited
jio

// jll : 2015-04-02 Jones Lang LaSalle Incorporated
jll

// jmp : 2015-03-26 Matrix IP LLC
jmp

// jnj : 2015-06-18 Johnson & Johnson Services, Inc.
jnj

// joburg : 2014-03-24 ZA Central Registry NPC trading as ZA Central Registry
joburg

// jot : 2014-12-18 Amazon Registry Services, Inc.
jot

// joy : 2014-12-18 Amazon Registry Services, Inc.
joy

// jpmorgan : 2015-04-30 JPMorgan Chase Bank, National Association
jpmorgan

// jprs : 2014-09-18 Japan Registry Services Co., Ltd.
jprs

// juegos : 2014-03-20 Uniregistry, Corp.
juegos

// juniper : 2015-07-30 JUNIPER NETWORKS, INC.
juniper

// kaufen : 2013-11-07 United TLD Holdco Ltd.
kaufen

// kddi : 2014-09-12 KDDI CORPORATION
kddi

// kerryhotels : 2015-04-30 Kerry Trading Co. Limited
kerryhotels

// kerrylogistics : 2015-04-09 Kerry Trading Co. Limited
kerrylogistics

// kerryproperties : 2015-04-09 Kerry Trading Co. Limited
kerryproperties

// kfh : 2014-12-04 Kuwait Finance House
kfh

// kia : 2015-07-09 KIA MOTORS CORPORATION
kia

// kim : 2013-09-23 Afilias plc
kim

// kinder : 2014-11-07 Ferrero Trading Lux S.A.
kinder

// kindle : 2015-06-25 Amazon Registry Services, Inc.
kindle

// kitchen : 2013-09-20 Binky Moon, LLC
kitchen

// kiwi : 2013-09-20 DOT KIWI LIMITED
kiwi

// koeln : 2014-01-09 punkt.wien GmbH
koeln

// komatsu : 2015-01-08 Komatsu Ltd.
komatsu

// kosher : 2015-08-20 Kosher Marketing Assets LLC
kosher

// kpmg : 2015-04-23 KPMG International Cooperative (KPMG International Genossenschaft)
kpmg

// kpn : 2015-01-08 Koninklijke KPN N.V.
kpn

// krd : 2013-12-05 KRG Department of Information Technology
krd

// kred : 2013-12-19 KredTLD Pty Ltd
kred

// kuokgroup : 2015-04-09 Kerry Trading Co. Limited
kuokgroup

// kyoto : 2014-11-07 Academic Institution: Kyoto Jyoho Gakuen
kyoto

// lacaixa : 2014-01-09 Fundación Bancaria Caixa d’Estalvis i Pensions de Barcelona, “la Caixa”
lacaixa

// ladbrokes : 2015-08-06 LADBROKES INTERNATIONAL PLC
ladbrokes

// lamborghini : 2015-06-04 Automobili Lamborghini S.p.A.
lamborghini

// lamer : 2015-10-01 The Estée Lauder Companies Inc.
lamer

// lancaster : 2015-02-12 LANCASTER
lancaster

// lancia : 2015-07-31 Fiat Chrysler Automobiles N.V.
lancia

// lancome : 2015-07-23 L'Oréal
lancome

// land : 2013-09-10 Binky Moon, LLC
land

// landrover : 2014-11-13 Jaguar Land Rover Ltd
landrover

// lanxess : 2015-07-30 LANXESS Corporation
lanxess

// lasalle : 2015-04-02 Jones Lang LaSalle Incorporated
lasalle

// lat : 2014-10-16 ECOM-LAC Federaciòn de Latinoamèrica y el Caribe para Internet y el Comercio Electrònico
lat

// latino : 2015-07-30 Dish DBS Corporation
latino

// latrobe : 2014-06-16 La Trobe University
latrobe

// law : 2015-01-22 Minds + Machines Group Limited
law

// lawyer : 2014-03-20 United TLD Holdco Ltd.
lawyer

// lds : 2014-03-20 IRI Domain Management, LLC ("Applicant")
lds

// lease : 2014-03-06 Binky Moon, LLC
lease

// leclerc : 2014-08-07 A.C.D. LEC Association des Centres Distributeurs Edouard Leclerc
leclerc

// lefrak : 2015-07-16 LeFrak Organization, Inc.
lefrak

// legal : 2014-10-16 Binky Moon, LLC
legal

// lego : 2015-07-16 LEGO Juris A/S
lego

// lexus : 2015-04-23 TOYOTA MOTOR CORPORATION
lexus

// lgbt : 2014-05-08 Afilias plc
lgbt

// liaison : 2014-10-02 Liaison Technologies, Incorporated
liaison

// lidl : 2014-09-18 Schwarz Domains und Services GmbH & Co. KG
lidl

// life : 2014-02-06 Binky Moon, LLC
life

// lifeinsurance : 2015-01-15 American Council of Life Insurers
lifeinsurance

// lifestyle : 2014-12-11 Lifestyle Domain Holdings, Inc.
lifestyle

// lighting : 2013-08-27 Binky Moon, LLC
lighting

// like : 2014-12-18 Amazon Registry Services, Inc.
like

// lilly : 2015-07-31 Eli Lilly and Company
lilly

// limited : 2014-03-06 Binky Moon, LLC
limited

// limo : 2013-10-17 Binky Moon, LLC
limo

// lincoln : 2014-11-13 Ford Motor Company
lincoln

// linde : 2014-12-04 Linde Aktiengesellschaft
linde

// link : 2013-11-14 Uniregistry, Corp.
link

// lipsy : 2015-06-25 Lipsy Ltd
lipsy

// live : 2014-12-04 United TLD Holdco Ltd.
live

// living : 2015-07-30 Lifestyle Domain Holdings, Inc.
living

// lixil : 2015-03-19 LIXIL Group Corporation
lixil

// llc : 2017-12-14 Afilias plc
llc

// loan : 2014-11-20 dot Loan Limited
loan

// loans : 2014-03-20 Binky Moon, LLC
loans

// locker : 2015-06-04 Dish DBS Corporation
locker

// locus : 2015-06-25 Locus Analytics LLC
locus

// loft : 2015-07-30 Annco, Inc.
loft

// lol : 2015-01-30 Uniregistry, Corp.
lol

// london : 2013-11-14 Dot London Domains Limited
london

// lotte : 2014-11-07 Lotte Holdings Co., Ltd.
lotte

// lotto : 2014-04-10 Afilias plc
lotto

// love : 2014-12-22 Merchant Law Group LLP
love

// lpl : 2015-07-30 LPL Holdings, Inc.
lpl

// lplfinancial : 2015-07-30 LPL Holdings, Inc.
lplfinancial

// ltd : 2014-09-25 Binky Moon, LLC
ltd

// ltda : 2014-04-17 InterNetX, Corp
ltda

// lundbeck : 2015-08-06 H. Lundbeck A/S
lundbeck

// lupin : 2014-11-07 LUPIN LIMITED
lupin

// luxe : 2014-01-09 Minds + Machines Group Limited
luxe

// luxury : 2013-10-17 Luxury Partners, LLC
luxury

// macys : 2015-07-31 Macys, Inc.
macys

// madrid : 2014-05-01 Comunidad de Madrid
madrid

// maif : 2014-10-02 Mutuelle Assurance Instituteur France (MAIF)
maif

// maison : 2013-12-05 Binky Moon, LLC
maison

// makeup : 2015-01-15 L'Oréal
makeup

// man : 2014-12-04 MAN SE
man

// management : 2013-11-07 Binky Moon, LLC
management

// mango : 2013-10-24 PUNTO FA S.L.
mango

// map : 2016-06-09 Charleston Road Registry Inc.
map

// market : 2014-03-06 United TLD Holdco Ltd.
market

// marketing : 2013-11-07 Binky Moon, LLC
marketing

// markets : 2014-12-11 Dotmarkets Registry Limited
markets

// marriott : 2014-10-09 Marriott Worldwide Corporation
marriott

// marshalls : 2015-07-16 The TJX Companies, Inc.
marshalls

// maserati : 2015-07-31 Fiat Chrysler Automobiles N.V.
maserati

// mattel : 2015-08-06 Mattel Sites, Inc.
mattel

// mba : 2015-04-02 Binky Moon, LLC
mba

// mckinsey : 2015-07-31 McKinsey Holdings, Inc.
mckinsey

// med : 2015-08-06 Medistry LLC
med

// media : 2014-03-06 Binky Moon, LLC
media

// meet : 2014-01-16 Charleston Road Registry Inc.
meet

// melbourne : 2014-05-29 The Crown in right of the State of Victoria, represented by its Department of State Development, Business and Innovation
melbourne

// meme : 2014-01-30 Charleston Road Registry Inc.
meme

// memorial : 2014-10-16 Dog Beach, LLC
memorial

// men : 2015-02-26 Exclusive Registry Limited
men

// menu : 2013-09-11 Wedding TLD2, LLC
menu

// merckmsd : 2016-07-14 MSD Registry Holdings, Inc.
merckmsd

// metlife : 2015-05-07 MetLife Services and Solutions, LLC
metlife

// miami : 2013-12-19 Minds + Machines Group Limited
miami

// microsoft : 2014-12-18 Microsoft Corporation
microsoft

// mini : 2014-01-09 Bayerische Motoren Werke Aktiengesellschaft
mini

// mint : 2015-07-30 Intuit Administrative Services, Inc.
mint

// mit : 2015-07-02 Massachusetts Institute of Technology
mit

// mitsubishi : 2015-07-23 Mitsubishi Corporation
mitsubishi

// mlb : 2015-05-21 MLB Advanced Media DH, LLC
mlb

// mls : 2015-04-23 The Canadian Real Estate Association
mls

// mma : 2014-11-07 MMA IARD
mma

// mobile : 2016-06-02 Dish DBS Corporation
mobile

// mobily : 2014-12-18 GreenTech Consultancy Company W.L.L.
mobily

// moda : 2013-11-07 United TLD Holdco Ltd.
moda

// moe : 2013-11-13 Interlink Co., Ltd.
moe

// moi : 2014-12-18 Amazon Registry Services, Inc.
moi

// mom : 2015-04-16 Uniregistry, Corp.
mom

// monash : 2013-09-30 Monash University
monash

// money : 2014-10-16 Binky Moon, LLC
money

// monster : 2015-09-11 Monster Worldwide, Inc.
monster

// mopar : 2015-07-30 FCA US LLC.
mopar

// mormon : 2013-12-05 IRI Domain Management, LLC ("Applicant")
mormon

// mortgage : 2014-03-20 United TLD Holdco Ltd.
mortgage

// moscow : 2013-12-19 Foundation for Assistance for Internet Technologies and Infrastructure Development (FAITID)
moscow

// moto : 2015-06-04 Motorola Trademark Holdings, LLC
moto

// motorcycles : 2014-01-09 DERMotorcycles, LLC
motorcycles

// mov : 2014-01-30 Charleston Road Registry Inc.
mov

// movie : 2015-02-05 Binky Moon, LLC
movie

// movistar : 2014-10-16 Telefónica S.A.
movistar

// msd : 2015-07-23 MSD Registry Holdings, Inc.
msd

// mtn : 2014-12-04 MTN Dubai Limited
mtn

// mtr : 2015-03-12 MTR Corporation Limited
mtr

// mutual : 2015-04-02 Northwestern Mutual MU TLD Registry, LLC
mutual

// nab : 2015-08-20 National Australia Bank Limited
nab

// nadex : 2014-12-11 Nadex Domains, Inc.
nadex

// nagoya : 2013-10-24 GMO Registry, Inc.
nagoya

// nationwide : 2015-07-23 Nationwide Mutual Insurance Company
nationwide

// natura : 2015-03-12 NATURA COSMÉTICOS S.A.
natura

// navy : 2014-03-06 United TLD Holdco Ltd.
navy

// nba : 2015-07-31 NBA REGISTRY, LLC
nba

// nec : 2015-01-08 NEC Corporation
nec

// netbank : 2014-06-26 COMMONWEALTH BANK OF AUSTRALIA
netbank

// netflix : 2015-06-18 Netflix, Inc.
netflix

// network : 2013-11-14 Binky Moon, LLC
network

// neustar : 2013-12-05 Registry Services, LLC
neustar

// new : 2014-01-30 Charleston Road Registry Inc.
new

// newholland : 2015-09-03 CNH Industrial N.V.
newholland

// news : 2014-12-18 United TLD Holdco Ltd.
news

// next : 2015-06-18 Next plc
next

// nextdirect : 2015-06-18 Next plc
nextdirect

// nexus : 2014-07-24 Charleston Road Registry Inc.
nexus

// nfl : 2015-07-23 NFL Reg Ops LLC
nfl

// ngo : 2014-03-06 Public Interest Registry
ngo

// nhk : 2014-02-13 Japan Broadcasting Corporation (NHK)
nhk

// nico : 2014-12-04 DWANGO Co., Ltd.
nico

// nike : 2015-07-23 NIKE, Inc.
nike

// nikon : 2015-05-21 NIKON CORPORATION
nikon

// ninja : 2013-11-07 United TLD Holdco Ltd.
ninja

// nissan : 2014-03-27 NISSAN MOTOR CO., LTD.
nissan

// nissay : 2015-10-29 Nippon Life Insurance Company
nissay

// nokia : 2015-01-08 Nokia Corporation
nokia

// northwesternmutual : 2015-06-18 Northwestern Mutual Registry, LLC
northwesternmutual

// norton : 2014-12-04 Symantec Corporation
norton

// now : 2015-06-25 Amazon Registry Services, Inc.
now

// nowruz : 2014-09-04 Asia Green IT System Bilgisayar San. ve Tic. Ltd. Sti.
nowruz

// nowtv : 2015-05-14 Starbucks (HK) Limited
nowtv

// nra : 2014-05-22 NRA Holdings Company, INC.
nra

// nrw : 2013-11-21 Minds + Machines GmbH
nrw

// ntt : 2014-10-31 NIPPON TELEGRAPH AND TELEPHONE CORPORATION
ntt

// nyc : 2014-01-23 The City of New York by and through the New York City Department of Information Technology & Telecommunications
nyc

// obi : 2014-09-25 OBI Group Holding SE & Co. KGaA
obi

// observer : 2015-04-30 Top Level Spectrum, Inc.
observer

// off : 2015-07-23 Johnson Shareholdings, Inc.
off

// office : 2015-03-12 Microsoft Corporation
office

// okinawa : 2013-12-05 BRregistry, Inc.
okinawa

// olayan : 2015-05-14 Crescent Holding GmbH
olayan

// olayangroup : 2015-05-14 Crescent Holding GmbH
olayangroup

// oldnavy : 2015-07-31 The Gap, Inc.
oldnavy

// ollo : 2015-06-04 Dish DBS Corporation
ollo

// omega : 2015-01-08 The Swatch Group Ltd
omega

// one : 2014-11-07 One.com A/S
one

// ong : 2014-03-06 Public Interest Registry
ong

// onl : 2013-09-16 I-Registry Ltd.
onl

// online : 2015-01-15 DotOnline Inc.
online

// onyourside : 2015-07-23 Nationwide Mutual Insurance Company
onyourside

// ooo : 2014-01-09 INFIBEAM INCORPORATION LIMITED
ooo

// open : 2015-07-31 American Express Travel Related Services Company, Inc.
open

// oracle : 2014-06-19 Oracle Corporation
oracle

// orange : 2015-03-12 Orange Brand Services Limited
orange

// organic : 2014-03-27 Afilias plc
organic

// origins : 2015-10-01 The Estée Lauder Companies Inc.
origins

// osaka : 2014-09-04 Osaka Registry Co., Ltd.
osaka

// otsuka : 2013-10-11 Otsuka Holdings Co., Ltd.
otsuka

// ott : 2015-06-04 Dish DBS Corporation
ott

// ovh : 2014-01-16 OVH SAS
ovh

// page : 2014-12-04 Charleston Road Registry Inc.
page

// panasonic : 2015-07-30 Panasonic Corporation
panasonic

// paris : 2014-01-30 City of Paris
paris

// pars : 2014-09-04 Asia Green IT System Bilgisayar San. ve Tic. Ltd. Sti.
pars

// partners : 2013-12-05 Binky Moon, LLC
partners

// parts : 2013-12-05 Binky Moon, LLC
parts

// party : 2014-09-11 Blue Sky Registry Limited
party

// passagens : 2015-03-05 Travel Reservations SRL
passagens

// pay : 2015-08-27 Amazon Registry Services, Inc.
pay

// pccw : 2015-05-14 PCCW Enterprises Limited
pccw

// pet : 2015-05-07 Afilias plc
pet

// pfizer : 2015-09-11 Pfizer Inc.
pfizer

// pharmacy : 2014-06-19 National Association of Boards of Pharmacy
pharmacy

// phd : 2016-07-28 Charleston Road Registry Inc.
phd

// philips : 2014-11-07 Koninklijke Philips N.V.
philips

// phone : 2016-06-02 Dish DBS Corporation
phone

// photo : 2013-11-14 Uniregistry, Corp.
photo

// photography : 2013-09-20 Binky Moon, LLC
photography

// photos : 2013-10-17 Binky Moon, LLC
photos

// physio : 2014-05-01 PhysBiz Pty Ltd
physio

// piaget : 2014-10-16 Richemont DNS Inc.
piaget

// pics : 2013-11-14 Uniregistry, Corp.
pics

// pictet : 2014-06-26 Pictet Europe S.A.
pictet

// pictures : 2014-03-06 Binky Moon, LLC
pictures

// pid : 2015-01-08 Top Level Spectrum, Inc.
pid

// pin : 2014-12-18 Amazon Registry Services, Inc.
pin

// ping : 2015-06-11 Ping Registry Provider, Inc.
ping

// pink : 2013-10-01 Afilias plc
pink

// pioneer : 2015-07-16 Pioneer Corporation
pioneer

// pizza : 2014-06-26 Binky Moon, LLC
pizza

// place : 2014-04-24 Binky Moon, LLC
place

// play : 2015-03-05 Charleston Road Registry Inc.
play

// playstation : 2015-07-02 Sony Computer Entertainment Inc.
playstation

// plumbing : 2013-09-10 Binky Moon, LLC
plumbing

// plus : 2015-02-05 Binky Moon, LLC
plus

// pnc : 2015-07-02 PNC Domain Co., LLC
pnc

// pohl : 2014-06-23 Deutsche Vermögensberatung Aktiengesellschaft DVAG
pohl

// poker : 2014-07-03 Afilias plc
poker

// politie : 2015-08-20 Politie Nederland
politie

// porn : 2014-10-16 ICM Registry PN LLC
porn

// pramerica : 2015-07-30 Prudential Financial, Inc.
pramerica

// praxi : 2013-12-05 Praxi S.p.A.
praxi

// press : 2014-04-03 DotPress Inc.
press

// prime : 2015-06-25 Amazon Registry Services, Inc.
prime

// prod : 2014-01-23 Charleston Road Registry Inc.
prod

// productions : 2013-12-05 Binky Moon, LLC
productions

// prof : 2014-07-24 Charleston Road Registry Inc.
prof

// progressive : 2015-07-23 Progressive Casualty Insurance Company
progressive

// promo : 2014-12-18 Afilias plc
promo

// properties : 2013-12-05 Binky Moon, LLC
properties

// property : 2014-05-22 Uniregistry, Corp.
property

// protection : 2015-04-23 XYZ.COM LLC
protection

// pru : 2015-07-30 Prudential Financial, Inc.
pru

// prudential : 2015-07-30 Prudential Financial, Inc.
prudential

// pub : 2013-12-12 United TLD Holdco Ltd.
pub

// pwc : 2015-10-29 PricewaterhouseCoopers LLP
pwc

// qpon : 2013-11-14 dotCOOL, Inc.
qpon

// quebec : 2013-12-19 PointQuébec Inc
quebec

// quest : 2015-03-26 Quest ION Limited
quest

// qvc : 2015-07-30 QVC, Inc.
qvc

// racing : 2014-12-04 Premier Registry Limited
racing

// radio : 2016-07-21 European Broadcasting Union (EBU)
radio

// raid : 2015-07-23 Johnson Shareholdings, Inc.
raid

// read : 2014-12-18 Amazon Registry Services, Inc.
read

// realestate : 2015-09-11 dotRealEstate LLC
realestate

// realtor : 2014-05-29 Real Estate Domains LLC
realtor

// realty : 2015-03-19 Fegistry, LLC
realty

// recipes : 2013-10-17 Binky Moon, LLC
recipes

// red : 2013-11-07 Afilias plc
red

// redstone : 2014-10-31 Redstone Haute Couture Co., Ltd.
redstone

// redumbrella : 2015-03-26 Travelers TLD, LLC
redumbrella

// rehab : 2014-03-06 United TLD Holdco Ltd.
rehab

// reise : 2014-03-13 Binky Moon, LLC
reise

// reisen : 2014-03-06 Binky Moon, LLC
reisen

// reit : 2014-09-04 National Association of Real Estate Investment Trusts, Inc.
reit

// reliance : 2015-04-02 Reliance Industries Limited
reliance

// ren : 2013-12-12 Beijing Qianxiang Wangjing Technology Development Co., Ltd.
ren

// rent : 2014-12-04 XYZ.COM LLC
rent

// rentals : 2013-12-05 Binky Moon, LLC
rentals

// repair : 2013-11-07 Binky Moon, LLC
repair

// report : 2013-12-05 Binky Moon, LLC
report

// republican : 2014-03-20 United TLD Holdco Ltd.
republican

// rest : 2013-12-19 Punto 2012 Sociedad Anonima Promotora de Inversion de Capital Variable
rest

// restaurant : 2014-07-03 Binky Moon, LLC
restaurant

// review : 2014-11-20 dot Review Limited
review

// reviews : 2013-09-13 United TLD Holdco Ltd.
reviews

// rexroth : 2015-06-18 Robert Bosch GMBH
rexroth

// rich : 2013-11-21 I-Registry Ltd.
rich

// richardli : 2015-05-14 Pacific Century Asset Management (HK) Limited
richardli

// ricoh : 2014-11-20 Ricoh Company, Ltd.
ricoh

// rightathome : 2015-07-23 Johnson Shareholdings, Inc.
rightathome

// ril : 2015-04-02 Reliance Industries Limited
ril

// rio : 2014-02-27 Empresa Municipal de Informática SA - IPLANRIO
rio

// rip : 2014-07-10 United TLD Holdco Ltd.
rip

// rmit : 2015-11-19 Royal Melbourne Institute of Technology
rmit

// rocher : 2014-12-18 Ferrero Trading Lux S.A.
rocher

// rocks : 2013-11-14 United TLD Holdco Ltd.
rocks

// rodeo : 2013-12-19 Minds + Machines Group Limited
rodeo

// rogers : 2015-08-06 Rogers Communications Canada Inc.
rogers

// room : 2014-12-18 Amazon Registry Services, Inc.
room

// rsvp : 2014-05-08 Charleston Road Registry Inc.
rsvp

// rugby : 2016-12-15 World Rugby Strategic Developments Limited
rugby

// ruhr : 2013-10-02 regiodot GmbH & Co. KG
ruhr

// run : 2015-03-19 Binky Moon, LLC
run

// rwe : 2015-04-02 RWE AG
rwe

// ryukyu : 2014-01-09 BRregistry, Inc.
ryukyu

// saarland : 2013-12-12 dotSaarland GmbH
saarland

// safe : 2014-12-18 Amazon Registry Services, Inc.
safe

// safety : 2015-01-08 Safety Registry Services, LLC.
safety

// sakura : 2014-12-18 SAKURA Internet Inc.
sakura

// sale : 2014-10-16 United TLD Holdco Ltd.
sale

// salon : 2014-12-11 Binky Moon, LLC
salon

// samsclub : 2015-07-31 Wal-Mart Stores, Inc.
samsclub

// samsung : 2014-04-03 SAMSUNG SDS CO., LTD
samsung

// sandvik : 2014-11-13 Sandvik AB
sandvik

// sandvikcoromant : 2014-11-07 Sandvik AB
sandvikcoromant

// sanofi : 2014-10-09 Sanofi
sanofi

// sap : 2014-03-27 SAP AG
sap

// sarl : 2014-07-03 Binky Moon, LLC
sarl

// sas : 2015-04-02 Research IP LLC
sas

// save : 2015-06-25 Amazon Registry Services, Inc.
save

// saxo : 2014-10-31 Saxo Bank A/S
saxo

// sbi : 2015-03-12 STATE BANK OF INDIA
sbi

// sbs : 2014-11-07 SPECIAL BROADCASTING SERVICE CORPORATION
sbs

// sca : 2014-03-13 SVENSKA CELLULOSA AKTIEBOLAGET SCA (publ)
sca

// scb : 2014-02-20 The Siam Commercial Bank Public Company Limited ("SCB")
scb

// schaeffler : 2015-08-06 Schaeffler Technologies AG & Co. KG
schaeffler

// schmidt : 2014-04-03 SALM S.A.S.
schmidt

// scholarships : 2014-04-24 Scholarships.com, LLC
scholarships

// school : 2014-12-18 Binky Moon, LLC
school

// schule : 2014-03-06 Binky Moon, LLC
schule

// schwarz : 2014-09-18 Schwarz Domains und Services GmbH & Co. KG
schwarz

// science : 2014-09-11 dot Science Limited
science

// scjohnson : 2015-07-23 Johnson Shareholdings, Inc.
scjohnson

// scor : 2014-10-31 SCOR SE
scor

// scot : 2014-01-23 Dot Scot Registry Limited
scot

// search : 2016-06-09 Charleston Road Registry Inc.
search

// seat : 2014-05-22 SEAT, S.A. (Sociedad Unipersonal)
seat

// secure : 2015-08-27 Amazon Registry Services, Inc.
secure

// security : 2015-05-14 XYZ.COM LLC
security

// seek : 2014-12-04 Seek Limited
seek

// select : 2015-10-08 iSelect Ltd
select

// sener : 2014-10-24 Sener Ingeniería y Sistemas, S.A.
sener

// services : 2014-02-27 Binky Moon, LLC
services

// ses : 2015-07-23 SES
ses

// seven : 2015-08-06 Seven West Media Ltd
seven

// sew : 2014-07-17 SEW-EURODRIVE GmbH & Co KG
sew

// sex : 2014-11-13 ICM Registry SX LLC
sex

// sexy : 2013-09-11 Uniregistry, Corp.
sexy

// sfr : 2015-08-13 Societe Francaise du Radiotelephone - SFR
sfr

// shangrila : 2015-09-03 Shangri‐La International Hotel Management Limited
shangrila

// sharp : 2014-05-01 Sharp Corporation
sharp

// shaw : 2015-04-23 Shaw Cablesystems G.P.
shaw

// shell : 2015-07-30 Shell Information Technology International Inc
shell

// shia : 2014-09-04 Asia Green IT System Bilgisayar San. ve Tic. Ltd. Sti.
shia

// shiksha : 2013-11-14 Afilias plc
shiksha

// shoes : 2013-10-02 Binky Moon, LLC
shoes

// shop : 2016-04-08 GMO Registry, Inc.
shop

// shopping : 2016-03-31 Binky Moon, LLC
shopping

// shouji : 2015-01-08 QIHOO 360 TECHNOLOGY CO. LTD.
shouji

// show : 2015-03-05 Binky Moon, LLC
show

// showtime : 2015-08-06 CBS Domains Inc.
showtime

// shriram : 2014-01-23 Shriram Capital Ltd.
shriram

// silk : 2015-06-25 Amazon Registry Services, Inc.
silk

// sina : 2015-03-12 Sina Corporation
sina

// singles : 2013-08-27 Binky Moon, LLC
singles

// site : 2015-01-15 DotSite Inc.
site

// ski : 2015-04-09 Afilias plc
ski

// skin : 2015-01-15 L'Oréal
skin

// sky : 2014-06-19 Sky International AG
sky

// skype : 2014-12-18 Microsoft Corporation
skype

// sling : 2015-07-30 Hughes Satellite Systems Corporation
sling

// smart : 2015-07-09 Smart Communications, Inc. (SMART)
smart

// smile : 2014-12-18 Amazon Registry Services, Inc.
smile

// sncf : 2015-02-19 Société Nationale des Chemins de fer Francais S N C F
sncf

// soccer : 2015-03-26 Binky Moon, LLC
soccer

// social : 2013-11-07 United TLD Holdco Ltd.
social

// softbank : 2015-07-02 SoftBank Corp.
softbank

// software : 2014-03-20 United TLD Holdco Ltd.
software

// sohu : 2013-12-19 Sohu.com Limited
sohu

// solar : 2013-11-07 Binky Moon, LLC
solar

// solutions : 2013-11-07 Binky Moon, LLC
solutions

// song : 2015-02-26 Amazon Registry Services, Inc.
song

// sony : 2015-01-08 Sony Corporation
sony

// soy : 2014-01-23 Charleston Road Registry Inc.
soy

// space : 2014-04-03 DotSpace Inc.
space

// spiegel : 2014-02-05 SPIEGEL-Verlag Rudolf Augstein GmbH & Co. KG
spiegel

// sport : 2017-11-16 Global Association of International Sports Federations (GAISF)
sport

// spot : 2015-02-26 Amazon Registry Services, Inc.
spot

// spreadbetting : 2014-12-11 Dotspreadbetting Registry Limited
spreadbetting

// srl : 2015-05-07 InterNetX, Corp
srl

// srt : 2015-07-30 FCA US LLC.
srt

// stada : 2014-11-13 STADA Arzneimittel AG
stada

// staples : 2015-07-30 Staples, Inc.
staples

// star : 2015-01-08 Star India Private Limited
star

// starhub : 2015-02-05 StarHub Ltd
starhub

// statebank : 2015-03-12 STATE BANK OF INDIA
statebank

// statefarm : 2015-07-30 State Farm Mutual Automobile Insurance Company
statefarm

// statoil : 2014-12-04 Statoil ASA
statoil

// stc : 2014-10-09 Saudi Telecom Company
stc

// stcgroup : 2014-10-09 Saudi Telecom Company
stcgroup

// stockholm : 2014-12-18 Stockholms kommun
stockholm

// storage : 2014-12-22 XYZ.COM LLC
storage

// store : 2015-04-09 DotStore Inc.
store

// stream : 2016-01-08 dot Stream Limited
stream

// studio : 2015-02-11 United TLD Holdco Ltd.
studio

// study : 2014-12-11 OPEN UNIVERSITIES AUSTRALIA PTY LTD
study

// style : 2014-12-04 Binky Moon, LLC
style

// sucks : 2014-12-22 Vox Populi Registry Ltd.
sucks

// supplies : 2013-12-19 Binky Moon, LLC
supplies

// supply : 2013-12-19 Binky Moon, LLC
supply

// support : 2013-10-24 Binky Moon, LLC
support

// surf : 2014-01-09 Minds + Machines Group Limited
surf

// surgery : 2014-03-20 Binky Moon, LLC
surgery

// suzuki : 2014-02-20 SUZUKI MOTOR CORPORATION
suzuki

// swatch : 2015-01-08 The Swatch Group Ltd
swatch

// swiftcover : 2015-07-23 Swiftcover Insurance Services Limited
swiftcover

// swiss : 2014-10-16 Swiss Confederation
swiss

// sydney : 2014-09-18 State of New South Wales, Department of Premier and Cabinet
sydney

// symantec : 2014-12-04 Symantec Corporation
symantec

// systems : 2013-11-07 Binky Moon, LLC
systems

// tab : 2014-12-04 Tabcorp Holdings Limited
tab

// taipei : 2014-07-10 Taipei City Government
taipei

// talk : 2015-04-09 Amazon Registry Services, Inc.
talk

// taobao : 2015-01-15 Alibaba Group Holding Limited
taobao

// target : 2015-07-31 Target Domain Holdings, LLC
target

// tatamotors : 2015-03-12 Tata Motors Ltd
tatamotors

// tatar : 2014-04-24 Limited Liability Company "Coordination Center of Regional Domain of Tatarstan Republic"
tatar

// tattoo : 2013-08-30 Uniregistry, Corp.
tattoo

// tax : 2014-03-20 Binky Moon, LLC
tax

// taxi : 2015-03-19 Binky Moon, LLC
taxi

// tci : 2014-09-12 Asia Green IT System Bilgisayar San. ve Tic. Ltd. Sti.
tci

// tdk : 2015-06-11 TDK Corporation
tdk

// team : 2015-03-05 Binky Moon, LLC
team

// tech : 2015-01-30 Personals TLD Inc.
tech

// technology : 2013-09-13 Binky Moon, LLC
technology

// telefonica : 2014-10-16 Telefónica S.A.
telefonica

// temasek : 2014-08-07 Temasek Holdings (Private) Limited
temasek

// tennis : 2014-12-04 Binky Moon, LLC
tennis

// teva : 2015-07-02 Teva Pharmaceutical Industries Limited
teva

// thd : 2015-04-02 Home Depot Product Authority, LLC
thd

// theater : 2015-03-19 Binky Moon, LLC
theater

// theatre : 2015-05-07 XYZ.COM LLC
theatre

// tiaa : 2015-07-23 Teachers Insurance and Annuity Association of America
tiaa

// tickets : 2015-02-05 Accent Media Limited
tickets

// tienda : 2013-11-14 Binky Moon, LLC
tienda

// tiffany : 2015-01-30 Tiffany and Company
tiffany

// tips : 2013-09-20 Binky Moon, LLC
tips

// tires : 2014-11-07 Binky Moon, LLC
tires

// tirol : 2014-04-24 punkt Tirol GmbH
tirol

// tjmaxx : 2015-07-16 The TJX Companies, Inc.
tjmaxx

// tjx : 2015-07-16 The TJX Companies, Inc.
tjx

// tkmaxx : 2015-07-16 The TJX Companies, Inc.
tkmaxx

// tmall : 2015-01-15 Alibaba Group Holding Limited
tmall

// today : 2013-09-20 Binky Moon, LLC
today

// tokyo : 2013-11-13 GMO Registry, Inc.
tokyo

// tools : 2013-11-21 Binky Moon, LLC
tools

// top : 2014-03-20 .TOP Registry
top

// toray : 2014-12-18 Toray Industries, Inc.
toray

// toshiba : 2014-04-10 TOSHIBA Corporation
toshiba

// total : 2015-08-06 Total SA
total

// tours : 2015-01-22 Binky Moon, LLC
tours

// town : 2014-03-06 Binky Moon, LLC
town

// toyota : 2015-04-23 TOYOTA MOTOR CORPORATION
toyota

// toys : 2014-03-06 Binky Moon, LLC
toys

// trade : 2014-01-23 Elite Registry Limited
trade

// trading : 2014-12-11 Dottrading Registry Limited
trading

// training : 2013-11-07 Binky Moon, LLC
training

// travel :  Dog Beach, LLC
travel

// travelchannel : 2015-07-02 Lifestyle Domain Holdings, Inc.
travelchannel

// travelers : 2015-03-26 Travelers TLD, LLC
travelers

// travelersinsurance : 2015-03-26 Travelers TLD, LLC
travelersinsurance

// trust : 2014-10-16 NCC Group Inc.
trust

// trv : 2015-03-26 Travelers TLD, LLC
trv

// tube : 2015-06-11 Latin American Telecom LLC
tube

// tui : 2014-07-03 TUI AG
tui

// tunes : 2015-02-26 Amazon Registry Services, Inc.
tunes

// tushu : 2014-12-18 Amazon Registry Services, Inc.
tushu

// tvs : 2015-02-19 T V SUNDRAM IYENGAR  & SONS LIMITED
tvs

// ubank : 2015-08-20 National Australia Bank Limited
ubank

// ubs : 2014-12-11 UBS AG
ubs

// uconnect : 2015-07-30 FCA US LLC.
uconnect

// unicom : 2015-10-15 China United Network Communications Corporation Limited
unicom

// university : 2014-03-06 Binky Moon, LLC
university

// uno : 2013-09-11 Dot Latin LLC
uno

// uol : 2014-05-01 UBN INTERNET LTDA.
uol

// ups : 2015-06-25 UPS Market Driver, Inc.
ups

// vacations : 2013-12-05 Binky Moon, LLC
vacations

// vana : 2014-12-11 Lifestyle Domain Holdings, Inc.
vana

// vanguard : 2015-09-03 The Vanguard Group, Inc.
vanguard

// vegas : 2014-01-16 Dot Vegas, Inc.
vegas

// ventures : 2013-08-27 Binky Moon, LLC
ventures

// verisign : 2015-08-13 VeriSign, Inc.
verisign

// versicherung : 2014-03-20 TLD-BOX Registrydienstleistungen GmbH
versicherung

// vet : 2014-03-06 United TLD Holdco Ltd.
vet

// viajes : 2013-10-17 Binky Moon, LLC
viajes

// video : 2014-10-16 United TLD Holdco Ltd.
video

// vig : 2015-05-14 VIENNA INSURANCE GROUP AG Wiener Versicherung Gruppe
vig

// viking : 2015-04-02 Viking River Cruises (Bermuda) Ltd.
viking

// villas : 2013-12-05 Binky Moon, LLC
villas

// vin : 2015-06-18 Binky Moon, LLC
vin

// vip : 2015-01-22 Minds + Machines Group Limited
vip

// virgin : 2014-09-25 Virgin Enterprises Limited
virgin

// visa : 2015-07-30 Visa Worldwide Pte. Limited
visa

// vision : 2013-12-05 Binky Moon, LLC
vision

// vistaprint : 2014-09-18 Vistaprint Limited
vistaprint

// viva : 2014-11-07 Saudi Telecom Company
viva

// vivo : 2015-07-31 Telefonica Brasil S.A.
vivo

// vlaanderen : 2014-02-06 DNS.be vzw
vlaanderen

// vodka : 2013-12-19 Minds + Machines Group Limited
vodka

// volkswagen : 2015-05-14 Volkswagen Group of America Inc.
volkswagen

// volvo : 2015-11-12 Volvo Holding Sverige Aktiebolag
volvo

// vote : 2013-11-21 Monolith Registry LLC
vote

// voting : 2013-11-13 Valuetainment Corp.
voting

// voto : 2013-11-21 Monolith Registry LLC
voto

// voyage : 2013-08-27 Binky Moon, LLC
voyage

// vuelos : 2015-03-05 Travel Reservations SRL
vuelos

// wales : 2014-05-08 Nominet UK
wales

// walmart : 2015-07-31 Wal-Mart Stores, Inc.
walmart

// walter : 2014-11-13 Sandvik AB
walter

// wang : 2013-10-24 Zodiac Wang Limited
wang

// wanggou : 2014-12-18 Amazon Registry Services, Inc.
wanggou

// warman : 2015-06-18 Weir Group IP Limited
warman

// watch : 2013-11-14 Binky Moon, LLC
watch

// watches : 2014-12-22 Richemont DNS Inc.
watches

// weather : 2015-01-08 International Business Machines Corporation
weather

// weatherchannel : 2015-03-12 International Business Machines Corporation
weatherchannel

// webcam : 2014-01-23 dot Webcam Limited
webcam

// weber : 2015-06-04 Saint-Gobain Weber SA
weber

// website : 2014-04-03 DotWebsite Inc.
website

// wed : 2013-10-01 Atgron, Inc.
wed

// wedding : 2014-04-24 Minds + Machines Group Limited
wedding

// weibo : 2015-03-05 Sina Corporation
weibo

// weir : 2015-01-29 Weir Group IP Limited
weir

// whoswho : 2014-02-20 Who's Who Registry
whoswho

// wien : 2013-10-28 punkt.wien GmbH
wien

// wiki : 2013-11-07 Top Level Design, LLC
wiki

// williamhill : 2014-03-13 William Hill Organization Limited
williamhill

// win : 2014-11-20 First Registry Limited
win

// windows : 2014-12-18 Microsoft Corporation
windows

// wine : 2015-06-18 Binky Moon, LLC
wine

// winners : 2015-07-16 The TJX Companies, Inc.
winners

// wme : 2014-02-13 William Morris Endeavor Entertainment, LLC
wme

// wolterskluwer : 2015-08-06 Wolters Kluwer N.V.
wolterskluwer

// woodside : 2015-07-09 Woodside Petroleum Limited
woodside

// work : 2013-12-19 Minds + Machines Group Limited
work

// works : 2013-11-14 Binky Moon, LLC
works

// world : 2014-06-12 Binky Moon, LLC
world

// wow : 2015-10-08 Amazon Registry Services, Inc.
wow

// wtc : 2013-12-19 World Trade Centers Association, Inc.
wtc

// wtf : 2014-03-06 Binky Moon, LLC
wtf

// xbox : 2014-12-18 Microsoft Corporation
xbox

// xerox : 2014-10-24 Xerox DNHC LLC
xerox

// xfinity : 2015-07-09 Comcast IP Holdings I, LLC
xfinity

// xihuan : 2015-01-08 QIHOO 360 TECHNOLOGY CO. LTD.
xihuan

// xin : 2014-12-11 Elegant Leader Limited
xin

// xn--11b4c3d : 2015-01-15 VeriSign Sarl
कॉम

// xn--1ck2e1b : 2015-02-26 Amazon Registry Services, Inc.
セール

// xn--1qqw23a : 2014-01-09 Guangzhou YU Wei Information Technology Co., Ltd.
佛山

// xn--30rr7y : 2014-06-12 Excellent First Limited
慈善

// xn--3bst00m : 2013-09-13 Eagle Horizon Limited
集团

// xn--3ds443g : 2013-09-08 TLD REGISTRY LIMITED
在线

// xn--3oq18vl8pn36a : 2015-07-02 Volkswagen (China) Investment Co., Ltd.
大众汽车

// xn--3pxu8k : 2015-01-15 VeriSign Sarl
点看

// xn--42c2d9a : 2015-01-15 VeriSign Sarl
คอม

// xn--45q11c : 2013-11-21 Zodiac Gemini Ltd
八卦

// xn--4gbrim : 2013-10-04 Suhub Electronic Establishment
موقع

// xn--55qw42g : 2013-11-08 China Organizational Name Administration Center
公益

// xn--55qx5d : 2013-11-14 China Internet Network Information Center (CNNIC)
公司

// xn--5su34j936bgsg : 2015-09-03 Shangri‐La International Hotel Management Limited
香格里拉

// xn--5tzm5g : 2014-12-22 Global Website TLD Asia Limited
网站

// xn--6frz82g : 2013-09-23 Afilias plc
移动

// xn--6qq986b3xl : 2013-09-13 Tycoon Treasure Limited
我爱你

// xn--80adxhks : 2013-12-19 Foundation for Assistance for Internet Technologies and Infrastructure Development (FAITID)
москва

// xn--80aqecdr1a : 2015-10-21 Pontificium Consilium de Comunicationibus Socialibus (PCCS) (Pontifical Council for Social Communication)
католик

// xn--80asehdb : 2013-07-14 CORE Association
онлайн

// xn--80aswg : 2013-07-14 CORE Association
сайт

// xn--8y0a063a : 2015-03-26 China United Network Communications Corporation Limited
联通

// xn--9dbq2a : 2015-01-15 VeriSign Sarl
קום

// xn--9et52u : 2014-06-12 RISE VICTORY LIMITED
时尚

// xn--9krt00a : 2015-03-12 Sina Corporation
微博

// xn--b4w605ferd : 2014-08-07 Temasek Holdings (Private) Limited
淡马锡

// xn--bck1b9a5dre4c : 2015-02-26 Amazon Registry Services, Inc.
ファッション

// xn--c1avg : 2013-11-14 Public Interest Registry
орг

// xn--c2br7g : 2015-01-15 VeriSign Sarl
नेट

// xn--cck2b3b : 2015-02-26 Amazon Registry Services, Inc.
ストア

// xn--cg4bki : 2013-09-27 SAMSUNG SDS CO., LTD
삼성

// xn--czr694b : 2014-01-16 Dot Trademark TLD Holding Company Limited
商标

// xn--czrs0t : 2013-12-19 Binky Moon, LLC
商店

// xn--czru2d : 2013-11-21 Zodiac Aquarius Limited
商城

// xn--d1acj3b : 2013-11-20 The Foundation for Network Initiatives “The Smart Internet”
дети

// xn--eckvdtc9d : 2014-12-18 Amazon Registry Services, Inc.
ポイント

// xn--efvy88h : 2014-08-22 Guangzhou YU Wei Information Technology Co., Ltd.
新闻

// xn--estv75g : 2015-02-19 Industrial and Commercial Bank of China Limited
工行

// xn--fct429k : 2015-04-09 Amazon Registry Services, Inc.
家電

// xn--fhbei : 2015-01-15 VeriSign Sarl
كوم

// xn--fiq228c5hs : 2013-09-08 TLD REGISTRY LIMITED
中文网

// xn--fiq64b : 2013-10-14 CITIC Group Corporation
中信

// xn--fjq720a : 2014-05-22 Binky Moon, LLC
娱乐

// xn--flw351e : 2014-07-31 Charleston Road Registry Inc.
谷歌

// xn--fzys8d69uvgm : 2015-05-14 PCCW Enterprises Limited
電訊盈科

// xn--g2xx48c : 2015-01-30 Minds + Machines Group Limited
购物

// xn--gckr3f0f : 2015-02-26 Amazon Registry Services, Inc.
クラウド

// xn--gk3at1e : 2015-10-08 Amazon Registry Services, Inc.
通販

// xn--hxt814e : 2014-05-15 Zodiac Taurus Limited
网店

// xn--i1b6b1a6a2e : 2013-11-14 Public Interest Registry
संगठन

// xn--imr513n : 2014-12-11 Dot Trademark TLD Holding Company Limited
餐厅

// xn--io0a7i : 2013-11-14 China Internet Network Information Center (CNNIC)
网络

// xn--j1aef : 2015-01-15 VeriSign Sarl
ком

// xn--jlq61u9w7b : 2015-01-08 Nokia Corporation
诺基亚

// xn--jvr189m : 2015-02-26 Amazon Registry Services, Inc.
食品

// xn--kcrx77d1x4a : 2014-11-07 Koninklijke Philips N.V.
飞利浦

// xn--kpu716f : 2014-12-22 Richemont DNS Inc.
手表

// xn--kput3i : 2014-02-13 Beijing RITT-Net Technology Development Co., Ltd
手机

// xn--mgba3a3ejt : 2014-11-20 Aramco Services Company
ارامكو

// xn--mgba7c0bbn0a : 2015-05-14 Crescent Holding GmbH
العليان

// xn--mgbaakc7dvf : 2015-09-03 Emirates Telecommunications Corporation (trading as Etisalat)
اتصالات

// xn--mgbab2bd : 2013-10-31 CORE Association
بازار

// xn--mgbb9fbpob : 2014-12-18 GreenTech Consultancy Company W.L.L.
موبايلي

// xn--mgbca7dzdo : 2015-07-30 Abu Dhabi Systems and Information Centre
ابوظبي

// xn--mgbi4ecexp : 2015-10-21 Pontificium Consilium de Comunicationibus Socialibus (PCCS) (Pontifical Council for Social Communication)
كاثوليك

// xn--mgbt3dhd : 2014-09-04 Asia Green IT System Bilgisayar San. ve Tic. Ltd. Sti.
همراه

// xn--mk1bu44c : 2015-01-15 VeriSign Sarl
닷컴

// xn--mxtq1m : 2014-03-06 Net-Chinese Co., Ltd.
政府

// xn--ngbc5azd : 2013-07-13 International Domain Registry Pty. Ltd.
شبكة

// xn--ngbe9e0a : 2014-12-04 Kuwait Finance House
بيتك

// xn--ngbrx : 2015-11-12 League of Arab States
عرب

// xn--nqv7f : 2013-11-14 Public Interest Registry
机构

// xn--nqv7fs00ema : 2013-11-14 Public Interest Registry
组织机构

// xn--nyqy26a : 2014-11-07 Stable Tone Limited
健康

// xn--otu796d : 2017-08-06 Dot Trademark TLD Holding Company Limited
招聘

// xn--p1acf : 2013-12-12 Rusnames Limited
рус

// xn--pbt977c : 2014-12-22 Richemont DNS Inc.
珠宝

// xn--pssy2u : 2015-01-15 VeriSign Sarl
大拿

// xn--q9jyb4c : 2013-09-17 Charleston Road Registry Inc.
みんな

// xn--qcka1pmc : 2014-07-31 Charleston Road Registry Inc.
グーグル

// xn--rhqv96g : 2013-09-11 Stable Tone Limited
世界

// xn--rovu88b : 2015-02-26 Amazon Registry Services, Inc.
書籍

// xn--ses554g : 2014-01-16 KNET Co., Ltd.
网址

// xn--t60b56a : 2015-01-15 VeriSign Sarl
닷넷

// xn--tckwe : 2015-01-15 VeriSign Sarl
コム

// xn--tiq49xqyj : 2015-10-21 Pontificium Consilium de Comunicationibus Socialibus (PCCS) (Pontifical Council for Social Communication)
天主教

// xn--unup4y : 2013-07-14 Binky Moon, LLC
游戏

// xn--vermgensberater-ctb : 2014-06-23 Deutsche Vermögensberatung Aktiengesellschaft DVAG
vermögensberater

// xn--vermgensberatung-pwb : 2014-06-23 Deutsche Vermögensberatung Aktiengesellschaft DVAG
vermögensberatung

// xn--vhquv : 2013-08-27 Binky Moon, LLC
企业

// xn--vuq861b : 2014-10-16 Beijing Tele-info Network Technology Co., Ltd.
信息

// xn--w4r85el8fhu5dnra : 2015-04-30 Kerry Trading Co. Limited
嘉里大酒店

// xn--w4rs40l : 2015-07-30 Kerry Trading Co. Limited
嘉里

// xn--xhq521b : 2013-11-14 Guangzhou YU Wei Information Technology Co., Ltd.
广东

// xn--zfr164b : 2013-11-08 China Organizational Name Administration Center
政务

// xyz : 2013-12-05 XYZ.COM LLC
xyz

// yachts : 2014-01-09 DERYachts, LLC
yachts

// yahoo : 2015-04-02 Yahoo! Domain Services Inc.
yahoo

// yamaxun : 2014-12-18 Amazon Registry Services, Inc.
yamaxun

// yandex : 2014-04-10 YANDEX, LLC
yandex

// yodobashi : 2014-11-20 YODOBASHI CAMERA CO.,LTD.
yodobashi

// yoga : 2014-05-29 Minds + Machines Group Limited
yoga

// yokohama : 2013-12-12 GMO Registry, Inc.
yokohama

// you : 2015-04-09 Amazon Registry Services, Inc.
you

// youtube : 2014-05-01 Charleston Road Registry Inc.
youtube

// yun : 2015-01-08 QIHOO 360 TECHNOLOGY CO. LTD.
yun

// zappos : 2015-06-25 Amazon Registry Services, Inc.
zappos

// zara : 2014-11-07 Industria de Diseño Textil, S.A. (INDITEX, S.A.)
zara

// zero : 2014-12-18 Amazon Registry Services, Inc.
zero

// zip : 2014-05-08 Charleston Road Registry Inc.
zip

// zippo : 2015-07-02 Zadco Company
zippo

// zone : 2013-11-14 Binky Moon, LLC
zone

// zuerich : 2014-11-07 Kanton Zürich (Canton of Zurich)
zuerich


// ===END ICANN DOMAINS===`
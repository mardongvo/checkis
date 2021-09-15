insert into staff(fio, position, phone, signer) select '_1','Директор филиала','',1 where not exists (select * from staff where fio='_1');
insert into staff(fio, position, phone, signer) select '_2','Заместитель директора филиала','',1 where not exists (select * from staff where fio='_2');

insert into staff(fio, position, phone, signer) select 's1','Главный специалист','',0 where not exists (select * from staff where fio='s1');
insert into staff(fio, position, phone, signer) select 's2','Главный специалист','',0 where not exists (select * from staff where fio='s2');

update staff set fio_dative = '_1', position_dative='Директору филиала' where fio='_1';

-- docs_req_flag_sicklist
insert into docs_list(flag, doc_title) values('docs_req_flag_sicklist', 'Документы, устанавливающие наличие трудовых отношений между страхователем и застрахованным лицом и подтверждающие страховой страж застрахованного лица (трудовая книжка, трудовой договор)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_sicklist', 'Заявление застрахованного лица о выплате пособия') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_sicklist', 'Заявление застрахованного лица о перерасчете (доплате) пособия (при необходимости)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_sicklist', 'Справка по перерасчету суммы (доплаты) пособия (при необходимости)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_sicklist', 'Листок нетрудоспособности (если представлен на бумажном носителе)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_sicklist', 'Сведения о сумме заработка застрахованного лица, из которого исчислено пособие (расчётные ведомости по заработной плате (по видам начислений и удержаний), лицевые счета или расчётные листки по заработной плате работников)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_sicklist', 'Справка (справки) о сумме заработка застрахованного лица, из которого должно быть исчислено пособие, с места (мест) работы (службы, иной деятельности) у другого страхователя (других страхователей)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_sicklist', 'В случае, если застрахованное лицо на момент наступления страхового случая занято у нескольких страхователей - справка (справки) с места работы (службы, иной деятельности) у другого страхователя (у других страхователей) о том, что назначение и выплата пособия этим страхователем не осуществляется (ч. 2.1. и 2.2. ст. 13 Закона № 255-ФЗ)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_sicklist', 'Документы, определяющие систему оплаты труда установленную страхователем («Положение об оплате труда», «Положение о премировании», «Коллективный договор», штатное расписание, приказы и распоряжения по предприятию, иные документы, определяющие систему оплаты труда страхователя и влияющие на исчисление заработка застрахованного лица)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_sicklist', 'Приказы и распоряжения по предприятию, иные документы, влияющие на выплату пособия (о простоях, отпусках, приеме, переводе, увольнении, режиме рабочего времени)') on conflict do nothing;

-- docs_req_flag_birth
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Документы, устанавливающие наличие трудовых отношений между страхователем и застрахованным лицом и подтверждающие страховой страж застрахованного лица (трудовая книжка, трудовой договор)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Заявление застрахованного лица о выплате пособия') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Заявление застрахованного лица о перерасчете (доплате) пособия (при необходимости)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Заявление застрахованного лица о предоставлении отпуска по беременности и родам') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Заявление застрахованного лица о замене календарных годов (при необходимости)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Приказ о предоставлении отпуска по беременности и родам') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Справка по перерасчёту суммы (доплаты) пособия (при необходимости)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Листок нетрудоспособности') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Сведения о сумме заработка застрахованного лица, из которого исчислено пособие (расчётные ведомости по заработной плате (по видам начислений и удержаний), лицевые счета или расчётные листки по заработной плате работников)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Справка (справки) о сумме заработка, из которого должно быть исчислено пособие, с места (мест) работы (службы, иной деятельности) у другого страхователя (других страхователей)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'В случае, если застрахованное лицо на момент наступления страхового случая занято у нескольких страхователей - справка (справки) с места работы (службы, иной деятельности) у другого страхователя (у других страхователей) о том, что назначение и выплата пособия этим страхователем не осуществляется (ч. 2.1. и 2.2. ст. 13 Закона № 255-ФЗ)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Документы, определяющие систему оплаты труда установленную страхователем («Положение об оплате труда», «Положение о премировании», «Коллективный договор», штатное расписание, приказы и распоряжения по предприятию, иные документы, определяющие систему оплаты труда страхователя и влияющие на исчисление заработка застрахованного лица)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_birth', 'Приказы и распоряжения по предприятию, иные документы, влияющие на выплату пособия (о простоях, отпусках, приеме, переводе, увольнении, режиме рабочего времени)') on conflict do nothing;
    
-- docs_req_flag_1_early
insert into docs_list(flag, doc_title) values('docs_req_flag_1_early', 'Документы, устанавливающие наличие трудовых отношений между страхователем и застрахованным лицом (трудовая книжка, трудовой договор, служебный контракт)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_1_early', 'Заявление застрахованного лица о выплате пособия') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_1_early', 'Справка о постановке на учёт в ранние сроки беременности') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_1_early', 'Листок нетрудоспособности, выданный  на период отпуска по беременности и родам (по основному месту работы)') on conflict do nothing;
    
-- docs_req_flag_1_birth
insert into docs_list(flag, doc_title) values('docs_req_flag_1_birth', 'Документы, устанавливающие наличие трудовых отношений между страхователем и застрахованным лицом (трудовая книжка, трудовой договор)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_1_birth', 'Заявление застрахованного лица о выплате пособия') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_1_birth', 'Справка о рождении ребенка (детей), выданную органами ЗАГС (иной документ в случае рождения ребёнка за пределами территории Российской Федерации)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_1_birth', 'Справка с места работы (службы, органа социальной защиты населения по месту жительства ребенка) другого родителя о том, что пособие не назначалось') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_1_birth', 'Выписка из решения об установлении над ребенком опеки (копия вступившего в законную силу решения суда об усыновлении, копия договора о передаче ребенка (детей) на воспитание в приемную семью) - для лица, заменяющего родителей (опекуна, усыновителя, приемного родителя)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_1_birth', 'Свидетельство о расторжении брака, - в случае, если брак между родителями расторгнут; документ, подтверждающий совместное проживание на территории Российской Федерации ребенка с одним из родителей, выданный организацией, уполномоченной на его выдачу') on conflict do nothing;
    
-- docs_req_flag_15_child
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Документы, устанавливающие наличие трудовых отношений между страхователем и застрахованным (трудовая книжка, трудовой договор)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Заявление застрахованного лица о выплате пособия') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Заявление застрахованного лица о перерасчете (доплате) пособия (при необходимости)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Справка по перерасчёту суммы (доплаты) пособия (при необходимости)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Заявление застрахованного лица о замене календарных лет (при необходимости)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Заявление застрахованного лица о предоставлении отпуска по уходу за ребенком') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Приказ о предоставлении застрахованному лицу отпуска по уходу за ребёнком') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Копия свидетельства о рождении (усыновлении) ребенка (детей), за которым осуществляется уход, и его копия либо выписка из решения об установлении над ребенком опеки') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Копия свидетельства о рождении (усыновлении) предыдущего ребёнка (детей) (в случае смерти предыдущего ребёнка предоставляется копия свидетельства о смерти)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Справка с места работы (службы, органа социальной защиты населения по месту жительства ребенка) другого родителя о том, что пособие не назначалось и не выплачивалось') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Сведения о сумме заработка застрахованного лица, из которого исчислено пособие (расчётные ведомости по заработной плате (по видам начислений и удержаний), лицевые счета или расчётные листки по заработной плате работников)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Справка (справки) о сумме заработка, из которого должно быть исчислено пособие, с места (мест) работы (службы, иной деятельности) у другого страхователя (других страхователей)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'В случае, если застрахованное лицо на момент наступления страхового случая занято у нескольких страхователей - справка (справки) с места работы (службы, иной деятельности) у другого страхователя (у других страхователей) о том, что назначение и выплата пособия этим страхователем не осуществляется (ч. 2.1. и 2.2. ст. 13 Закона № 255-ФЗ)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Документы, определяющие систему оплаты труда, установленную страхователем («Положение об оплате труда», «Положение о премировании», «Коллективный договор», штатное расписание, приказы и распоряжения по предприятию, иные документы, определяющие систему оплаты труда страхователя и влияющие на исчисление заработка застрахованного лица)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_15_child', 'Приказы и распоряжения по предприятию, иные документы, влияющие на выплату пособия (о простоях, отпусках, приеме, переводе, увольнении, режиме рабочего времени)') on conflict do nothing;
    
-- docs_req_flag_accident
insert into docs_list(flag, doc_title) values('docs_req_flag_accident', 'Заявление застрахованного лица о выплате пособия') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_accident', 'Заявление застрахованного лица о перерасчете (доплате) пособия (при необходимости)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_accident', 'Документы, устанавливающие наличие трудовых отношений между страхователем и застрахованным лицом (трудовая книжка, трудовой договор)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_accident', 'Листок нетрудоспособности') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_accident', 'Справка по перерасчёту суммы (доплаты) пособия (при необходимости)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_accident', 'Сведения о сумме заработка застрахованного лица, из которого исчислено пособие (расчётные ведомости по заработной плате (по видам начислений и удержаний), лицевые счета или расчётные листки по заработной плате работников), документы, определяющие систему оплаты труда, установленную страхователем («Положение об оплате труда», «Положение о премировании», «Коллективный договор», штатное расписание, приказы и распоряжения по предприятию, иные документы, определяющие систему оплаты труда страхователя и влияющие на исчисление заработка застрахованного лица') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_accident', 'Акт о несчастном случае (форма Н-1) либо акт о случае профессионального заболевания') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_accident', 'Справка (справки) о сумме заработка, из которого должно быть исчислено пособие, с места (мест) работы (службы, иной деятельности) у другого страхователя (других страхователей)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_accident', 'Приказы и распоряжения по предприятию, иные документы, влияющие на выплату пособия (о простоях, отпусках, приеме, переводе, увольнении, режиме рабочего времени)') on conflict do nothing;
    
-- docs_req_flag_vacation
  
insert into docs_list(flag, doc_title) values('docs_req_flag_vacation', 'Заявление застрахованного лица об оплате отпуска') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_vacation', 'Заявление застрахованного лица о перерасчете (доплате) пособия (при необходимости)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_vacation', 'Приказ о предоставлении дополнительного отпуска работнику') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_vacation', 'Документы, устанавливающие наличие трудовых отношений между страхователем и застрахованным лицом (трудовая книжка, трудовой договор)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_vacation', 'Справка-расчет размера оплаты отпуска (сверх ежегодного оплачиваемого отпуска) на весь период лечения и проезда к месту лечения и обратно') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_vacation', 'Сведения о сумме заработка застрахованного лица, из которого исчислено пособие (расчётные ведомости по заработной плате (по видам начислений и удержаний), лицевые счета или расчётные листки по заработной плате работников)
табели учёта рабочего времени за расчётный период') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_vacation', 'Документы, определяющие систему оплаты труда установленную страхователем («Положение об оплате труда», «Положение о премировании», «Коллективный договор», штатное расписание, приказы и распоряжения по предприятию, иные документы, определяющие систему оплаты труда страхователя и влияющие на исчисление заработка застрахованного лица)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_vacation', 'Документы(билеты), свидетельствующие об оплате проезда к месту лечения и обратно') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_vacation', 'Документы,свидетельствующие об затратах на проезд к месту лечения и обратно') on conflict do nothing;
    
-- docs_req_flag_4days
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Заявление о возмещении расходов') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Документы, устанавливающие наличие трудовых отношений между страхователем и застрахованным лицом (трудовая книжка, трудовой договор)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Заверенная копия приказа о предоставлении дополнительных выходных дней одному из родителей (опекуну, попечителю) для ухода за детьми-инвалидами') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Заявление застрахованного лица о предоставлении дополнительного выходного дня (дней)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Сведения о начисленных страховых взносах в государственные внебюджетные фонды при оплате 4 выходных дополнительных дней одному из родителей (опекуну, попечителю) для ухода за детьми-инвалидами') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Справка, подтверждающая факт установления инвалидности, выданная бюро (главным бюро, Федеральным бюро) медико-социальной экспертизы') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Свидетельство о рождении (усыновлении) ребенка либо документ, подтверждающий установление опеки, попечительства над ребенком-инвалидом') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Документы, подтверждающие место жительства (пребывания или фактического проживания)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Справка с места работы другого родителя (опекуна, попечителя) о том, что на момент обращения дополнительные оплачиваемые выходные дни в этом же календарном месяце им не использованы или использованы частично, либо справка с места работы другого родителя (опекуна, попечителя) о том, что от этого родителя (опекуна, попечителя) не поступало заявления о предоставлении ему в этом же календарном месяце дополнительных оплачиваемых выходных дней') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Если один из родителей (опекунов, попечителей) не состоит в трудовых отношениях либо является индивидуальным предпринимателем, адвокатом, нотариусом, занимающимся частной практикой, или иным лицом, занимающимся в установленном законодательством Российской Федерации порядке частной практикой, членом зарегистрированных в установленном порядке семейных (родовых) общин коренных малочисленных народов Севера, Сибири и Дальнего Востока Российской Федерации, родитель (опекун, попечитель), состоящий в трудовых отношениях, представляет работодателю документы (их копии), подтверждающие указанные факты, при каждом обращении с заявлением') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Сведения о сумме заработка застрахованного лица, из которого исчислено пособие (расчётные ведомости по заработной плате (по видам начислений и удержаний), лицевые счета или расчётные листки по заработной плате работников)
табели учета  рабочего времени за расчётный период') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_4days', 'Документы, определяющие систему оплаты труда установленную страхователем («Положение об оплате труда», «Положение о премировании», «Коллективный договор», штатное расписание, приказы и распоряжения по предприятию, иные документы, определяющие систему оплаты труда страхователя и влияющие на исчисление заработка застрахованного лица)') on conflict do nothing;
    
-- docs_req_flag_burial_soc
insert into docs_list(flag, doc_title) values('docs_req_flag_burial_soc', 'Заявление о возмещении указанных расходов') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_burial_soc', 'Документы, устанавливающие наличие трудовых отношений между страхователем и застрахованным лицом на момент его смерти (выписка из трудовой книжки, личная карта работника (форма Т-2)) либо на момент смерти несовершеннолетнего члена семьи застрахованного лица (трудовая книжка, трудовой договор)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_burial_soc', 'Справка о смерти застрахованного лица (либо несовершеннолетнего члена семьи застрахованного лица)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_burial_soc', 'Заявление супруга, близкого родственника, иного родственника, законного представителя умершего или иного лица, взявшего на себя обязанность осуществить погребение умершего') on conflict do nothing;

-- docs_req_flag_burial_spec
insert into docs_list(flag, doc_title) values('docs_req_flag_burial_spec', 'Заявление специализированной службы по вопросам похоронного дела о возмещении стоимости гарантированного перечня услуг по погребению') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_burial_spec', 'Документы, устанавливающие наличие трудовых отношений между страхователем и застрахованным лицом на момент его смерти (выписка из трудовой книжки, личная карта работника (форма Т-2), либо на момент смерти несовершеннолетнего члена семьи застрахованного лица (трудовая книжка, трудовой договор, служебный контракт)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_burial_spec', 'Справка о смерти застрахованного лица (либо несовершеннолетнего члена семьи застрахованного лица)') on conflict do nothing;
insert into docs_list(flag, doc_title) values('docs_req_flag_burial_spec', 'Счёт на оплату указанных услуг') on conflict do nothing;



var targetHost = 'http://192.168.1.220:8088'

var selectedBaseType = 'none'
var selectedOperation = 'none'

$(function(){
    $( document ).ajaxStart(function() {
        Messenger().post( "Triggered ajaxStart handler." );
    });

    $('#search-form').submit(function(){
        var playerId = getPlayerId()
        if(playerId == 0){
            return
        }

        showRoleInfo(playerId)
        return false;
    })

    $("#delete_order").click(function(){
        orderList = []
        $("#gmOrder tbody input:checked").each(function(i,o){
            orderList.push($(o).val())
        })

        if(orderList.length == 0){
            Messenger().post({
                message: 'Please select order to delete.',
                type: 'error',
                showCloseButton: true
            });
            return
        }

        var playerId = getPlayerId()
        if(playerId == 0){
            return
        }

        $.ajax({
            url:targetHost+"/gmcommand/del?RoleUid="+playerId+"&OrderUid="+orderList.join(',',orderList),
            dataType:'json',
            timeout:5000})
            .done(function(d){
                if(d.retCode != 'Del GmOrderList Success'){
                    Messenger().post({
                        message: d.retCode,
                        type: 'error',
                        showCloseButton: true
                    });
                    return
                }
                //reload
                showRoleInfo(playerId)
            })
            .error(function(XMLHttpRequest, textStatus, errorThrown){
                console.log(XMLHttpRequest)
                console.log(textStatus)
                console.log(errorThrown)
            })
    })

    $('#edit').on('click', '.btn-success', function(){
        if($(this).children('.fa-plus').length >0){
            var option = $(this).parents('.form-group').prop("outerHTML")
            option = $(option).find('.fa-plus').removeClass('fa-plus').addClass('fa-minus').end().prop("outerHTML")
            $(this).parents('.form-group').after(option)
            //fix selectpicker bug
            $('select[class*=optionOption]').selectpicker('refresh');
            $('div.optionOption').next('div.optionOption').remove();

        }else{
            $(this).parents('.form-group').remove()
        }
    })

    BaseTypeOption = ''
    for(i in CommandBaseType){
        BaseTypeOption += "<option value='"+i+"'>"+CommandBaseType[i]+"</option>"
    }
    $('#selectBaseType').append(BaseTypeOption)

    OperationOption = ''
    for(i in CommandOperation){
        OperationOption += "<option value='"+i+"'>"+CommandOperation[i]+"</option>"
    }
    $('#selectOperation').append(OperationOption)

    OptionOption = ''
    for(i in CommandOption){
        OptionOption += "<option value='"+i+"' name='"+CommandOption[i]+"'>"+CommandOption[i]+"</option>"
    }
    $('.optionOption').append(OptionOption)

    $('#edit #selectBaseType').change(function(d){
        selectedBaseType  = $(this).val()
        if(selectedBaseType == 'none'){
            //reset
        }else{
            //reset operation
            $('#selectOperation option').prop('class','')
            $('#selectOperation').selectpicker('refresh')
            $('#selectOperation').selectpicker('val','none')

            //reset option
            $('select[class*=optionOption] option').prop('class','')
            $('select[class*=optionOption]').selectpicker('refresh')
            $('select[class*=optionOption]').selectpicker('val','none')

            //reset option type
            $('select[class*=optionType] option').prop('class','')
            $('select[class*=optionType]').selectpicker('refresh')
            $('select[class*=optionType]').selectpicker('val','0')

            switch (selectedBaseType){
                case "0"://role
                    var temp = ["Exp","Order","Spirit","Gold","FreeGold","Blood","Soul","Stone"]
                    $('#selectOperation option:contains("New")').prop('class','hide')
                    for(var i in temp){
                        $('select[class*=optionOption] option[name="'+temp[i]+'"]').prop('class','show')
                    }
                    break;
                case "1"://king
                    var temp = ["Lv","SkillLv"]
                    $('#selectOperation option:contains("New")').prop('class','hide')
                    for(var i in temp){
                        $('select[class*=optionOption] option[name="'+temp[i]+'"]').prop('class','show')
                    }
                    break;
                case "2"://Item
                    var temp = ["Num"]
                    for(var i in temp){
                        $('select[class*=optionOption] option[name="'+temp[i]+'"]').prop('class','show')
                    }
                    break;
                case "3"://Hero
                    var temp = ["Lv","Rank","SkillLv"]
                    for(var i in temp){
                        $('select[class*=optionOption] option[name="'+temp[i]+'"]').prop('class','show')
                    }
                    break;
                case "4"://Soldier
                    var temp = ["Lv","Num","SkillLv"]
                    $('#selectOperation option:contains("New")').prop('class','hide')
                    for(var i in temp){
                        $('select[class*=optionOption] option[name="'+temp[i]+'"]').prop('class','show')
                    }
                    break;
                case "5"://Building
                    var temp = ["Lv"]
                    console.log(temp)
                    $('#selectOperation option:contains("New")').prop('class','hide')
                    for(var i in temp){
                        $('select[class*=optionOption] option[name="'+temp[i]+'"]').prop('class','show')
                    }
                    break;
            }

            $('#selectOperation').selectpicker('refresh')
            $('select[class*=optionOption] option[class!=show]').prop('class','hide')
            $('select[class*=optionOption] option[value="none"]').prop('class','show')

            $('select[class*=optionOption]').selectpicker('refresh')
            $('select[class*=optionOption]').selectpicker('val','none')
            $('div.optionType').next('div.optionType').remove();
            refreshSchemeInput()
        }
    })

    $('#edit').on('change','#selectOperation',function(){
        selectedOperation = $(this).val()
        if(selectedOperation == 'none'){

        }else{
            switch (selectedOperation){
                case '0'://New
                    //hero
                    if (selectedBaseType == '3'){
                        $('select[class*=optionOption] option[name="SkillLv"]').prop('class','hide')
                    }
                    break;
                case '1'://Fix
                    if (selectedBaseType == '3'){
                        $('select[class*=optionOption] option[name="SkillLv"]').prop('class','show')
                    }
                    break
            }
            $('select[class*=optionOption]').selectpicker('refresh')

            $('select[class*=optionOption]').selectpicker('val','none')
            refreshSchemeInput()
            $('select[class*=optionType] option').prop('class','')
            refreshOptionType()
        }
    })

    $('#edit').on('change','.optionOption',function(){
        var selectedOption = $(this).val()
        $('select[class*=optionType] option').prop('class','')
        var optionGroup = $(this).parents('.form-group')
        //role
        if (selectedBaseType=='0'){
            if(selectedOption== '1'){
                optionGroup.find('select[class*=optionType] option:contains("Edit")').prop('class','hide')
                optionGroup.find('select[class*=optionType]').selectpicker('val',1)
            }else{
                optionGroup.find('select[class*=optionType] option:contains("Edit")').prop('class','')
            }
            //king
        }else if(selectedBaseType == '1'){
            //lv
            if(selectedOption == '0'){
                optionGroup.find('select[class*=optionType] option:contains("Edit")').prop('class','hide')
                optionGroup.find('select[class*=optionType]').selectpicker('val',1)
            }else{
                optionGroup.find('select[class*=optionType] option:contains("Add")').prop('class','hide')
                optionGroup.find('select[class*=optionType]').selectpicker('val',0)
            }
        }

        refreshOptionType()
    })

    $('#edit #save').click(function(){
        var playerId = $('#tooltip-enabled').val()
        if (isNaN(Number(playerId))){
            Messenger().post({
                message: 'Please input player id.',
                type: 'error',
                showCloseButton: true
            });
            return
        }
        if(playerId==0){
            return
        }

        $('#search-form input').val(playerId)

        var SchemeId = $('#UidOrSchemeId').val()
        if(SchemeId == ''){
            SchemeId = playerId
        }

        var BaseType = $('#selectBaseType').val()
        var Operation = $('#selectOperation').val()

        if(BaseType == 'none'|| Operation == 'none'){
            Messenger().post({
                message: 'All Options need to be selected.',
                type: 'error',
                showCloseButton: true
            });
            return
        }

        var optionValue=[]
        var optionType =[]
        var option = []
        $('input[class*="option"]').each(function(i,o){
            var v = $(o).val()
            if(v == ''){
                Messenger().post({
                    message: 'All values must be number.',
                    type: 'error',
                    showCloseButton: true
                });
                return
            }

            if(v.indexOf(',') != -1){
                var temp = v.split(',')
                optionValue.push(temp[0]*1000+temp[1]*1)
            }else{
                optionValue.push(v)
            }
        })

        $('.optionType button .filter-option').each(function(i,o){
            var v = $(o).html()
            optionType.push(v)
        })

        $('.optionOption option:selected').each(function(i,o){
            var v = $(o).val()
            if(v == 'none'){
                Messenger().post({
                    message: 'All Options need to be selected.',
                    type: 'error',
                    showCloseButton: true
                });
                return
            }
            option.push(CommandOption[v])
        })

        //options
        options = []
        for(i=0; i<option.length; i++){
            options.push(option[i]+'='+optionType[i]+','+optionValue[i])
        }

        if ($('#UidOrSchemeId:visible').length > 0 &&$('#UidOrSchemeId:visible').val() == ''){
            Messenger().post({
                message: 'All values must be number..',
                type: 'error',
                showCloseButton: true
            });
            return
        }

        $.ajax({
            url:targetHost+"/gmcommand/process?RoleUid="+playerId+"&UidOrSchemeId="+SchemeId+"&Module="+CommandBaseType[BaseType]+"&Property="+CommandOperation[Operation]+'&'+options.join('&'),
            dataType:'json',
            timeout:5000})
            .done(function(d){
                if(d.retCode != 'Create GmOrder Success'){
                    Messenger().post({
                        message: d.retCode,
                        type: 'error',
                        showCloseButton: true
                    });
                }else{
                    Messenger().post({
                        message: d.retCode,
                        type: 'success',
                        showCloseButton: true
                    });
                    showRoleInfo(playerId)
                }
            })
    })
})

var CommandBaseType = {
    0 : "Role",
    1 : "King",
    2 : "Item",
    3 : "Hero",
    4 : "Soldier",
    5 : "Building"
}

var CommandOperation = {
    0:"New",
    1:"Fix"
}

var CommandOption = {
    0: "Lv",
    1: "Exp",
    2: "Order",
    3: "Spirit",
    4: "Soul",
    5: "Blood",
    6: "Stone",
    7: "Gold",
    8: "FreeGold",
    9: "SkillLv",
    10:"Num",
    11:"Stage",
    12:"Rank"
}

var CommandOptionType = {
    0:"Edit",
    1:"Add"
}

var OrderStatus = {
    0:"NoProcess",
    1:"Success",
    2:"Fail"
}

function refreshOptionType(){
    if (selectedBaseType == '2'){
        if(selectedOperation == '0'){
            $('select[class*=optionType] option:contains("Add")').prop('class','hide')
            $('select[class*=optionType]').selectpicker('val',0)
        }
    }else if(selectedBaseType == '3'){
        $('select[class*=optionType] option:contains("Add")').prop('class','hide')
        $('select[class*=optionType]').selectpicker('val', 0)
    }else if(selectedBaseType == '4' || selectedBaseType == '5'){
        $('select[class*=optionType] option:contains("Add")').prop('class','hide')
        $('select[class*=optionType]').selectpicker('val', 0)
    }

    $('select[class*=optionType]').selectpicker('refresh')
    $('div.optionType').next('div.optionType').remove()
}

function refreshSchemeInput(){
    var flag = 'show'
    var schemeIdText = 'SchemeId'
    switch (selectedBaseType){
        case "0":
            flag = 'hide'
            break;
        case "1":
            flag = 'hide'
            break;
        case "2":
        case "3":
            flag = 'show'
            //New
            if(selectedOperation != '0'){
                schemeIdText = 'Uid'
            }
            break;
        case "4":
            flag = 'show'
            break;
        case "5":
            flag = "show"
            schemeIdText = 'Uid'
            break;
    }
    $('#schemeId').parent().prop('class', flag)
    $('#schemeId').html(schemeIdText)
}


function getPlayerId(){
    var playerId = $('#search-form input').val()
    if (isNaN(Number(playerId))){
        Messenger().post({
            message: 'Please input number.',
            type: 'error',
            showCloseButton: true
        });
        return 0
    }

    return playerId;
}

function showRoleInfo(playerId) {
    $.ajax({url:targetHost+"/gmcommand/query?RoleUid="+playerId,dataType:'json',timeout:5000})
        .done(showAll)
        .error(function(XMLHttpRequest, textStatus, errorThrown){
            Messenger().post({
                message: textStatus,
                type: 'error',
                showCloseButton: true
            });
        })
}

function showAll(d){
    if (d.retCode != "Query RoleAllInfo Success") {
        Messenger().post({
            message: d.retCode,
            type: 'error',
            showCloseButton: true
        });
        return
    }

    Messenger().post({
        message: d.retCode,
        type: 'success',
        showCloseButton: true
    });
    $('#tooltip-enabled').val($('#search-form input').val())

    role = d.content.role
    html = '<tr>\
        <th>' + role.uid + '</th>\
        <th>' + role.nickname + '</th>\
        <th>' + role.lv + '</th>\
        <th>' + role.exp + '</th>\
        <th>' + role.order + '</th>\
        <th>' + role.spirit + '</th>\
        <th>' + role.gold + '</th>\
        <th>' + role.free_gold + '</th>\
        <th>' + role.stone + '</th>\
        <th>' + role.soul + '</th>\
        <th>' + role.blood + '</th>\
        </tr>'
    $('#role tbody').html(html)

    king = d.content.king
    skills = ''
    for (i in king.king_skills) {
        skill = king.king_skills[i]
        if (skill.skill_id == 0){
            continue;
        }
        skills += '<th>'+KingSkillName[skill.skill_id]+'id:' + skill.skill_id + ' lv:' + skill.lv + '</th>';
    }

    html = '<tr>\
        <th>' + king.king_lv + '</th>\
        '+skills+'\
        </tr>'
    $('#king tbody').html(html)

    items = d.content.item
    html = ''
    for (o in items) {
        item = items[o]
        html += '<tr>\
        <th>' + item.uid + '</th>\
        <th>' +ItemName[item.scheme_id]+ ':'+item.scheme_id + '</th>\
        <th>' + item.num + '</th>\
        </tr>'
    }
    $('#item tbody').html(html)

    heros = d.content.hero
    html = ''
    for (o in heros) {
        hero = heros[o]
        skills = ''
        var temp = 0
        for (i in hero.skill_list) {
            skill = hero.skill_list[i]
            skills += '<th>'+SkillName[skill.skill_id]+'id:' + skill.skill_id + ' lv:' + skill.skill_lv + '</th>';
            temp++

            if (temp == Object.keys(hero.skill_list).length-1){
                break;
            }

        }

        for (i = 0; i < 4 - temp; i++) {
            skills += '<th></th>';
        }

        html += '<tr>\
        <th>' + hero.uid + '</th>\
        <th>' + HeroName[hero.scheme_id] +':'+ hero.scheme_id + '</th>\
        <th>' + hero.lv + '</th>\
        <th>' + hero.stage + '</th>\
        <th>' + hero.rank + '</th>'
            + skills +
            '</tr>'
    }
    $('#hero tbody').html(html)

    soldiersCamp = d.content.soldier
    html = ''
    for (o in soldiersCamp) {
        soldier = soldiersCamp[o]
        skills = ''
        for (i in soldier.skillLevel) {
            lv = soldier.skillLevel[i]
            skills += '<th>'+SkillName[i]+'id:' + i + ' lv:' + lv + '</th>';
            break;
        }

        html += '<tr>\
        <th>' +SoldierName[soldier.schemeId]+':'+ soldier.schemeId + '</th>\
        <th>' + soldier.num + '</th>\
        <th>' + soldier.level + '</th>\
        <th>' + soldier.stage + '</th>'
            + skills +
            '</tr>'
    }
    $('#soldier tbody').html(html)

    buildings = d.content.building
    html = ''
    for (o in buildings) {
        building = buildings[o]

        var collect_timestamp = ''
        if(building.collect_timestamp != 0){
            temp = new Date(parseInt(building.collect_timestamp)*1000)
            collect_timestamp = temp.toISOString()
        }

        html += '<tr>\
        <th>' + building.uid + '</th>\
        <th>' +BuildingName[building.scheme_id]+':'+ building.scheme_id + '</th>\
        <th>' + building.lv + '</th>\
        </tr>'
    }
    $('#building tbody').html(html)

    gmOrders = d.content.gmOrder

    html = ''
    var checkList = 1;
    for (o in gmOrders) {
        gmOrder = gmOrders[o]

        content = ''
        for (i in gmOrder.content) {
            ii = gmOrder.content[i]
            content +=   CommandOptionType[ii.operation]+':'+CommandOption[ii.option] + ':' + ii.value +'<br>';
        }

        status = 'default'
        if (gmOrder.orderStatus == 0){
            status = 'default'
        }else if (gmOrder.orderStatus == 1){
            status = 'success'
        }else if (gmOrder.orderStatus == 2){
            status = 'danger'
        }

        var time = new Date(parseInt(gmOrder.orderId))
        checkList++
        html += '<tr>\
            <th>\
                <div class="checkbox">\
                <input id="checkbox'+checkList+'" type="checkbox" value="'+gmOrder.orderId+'">\
                <label for="checkbox'+checkList+'"></label>\
                </div>\
                </th>\
        <th>' + time.toISOString() + '</th>\
        <th>' + CommandBaseType[gmOrder.commandModule] + '</th>\
        <th>' + CommandOperation[gmOrder.commandProperty] + '</th>\
        <th>' + gmOrder.uidOrSchemeId + '</th>\
        <th>'+content+'</th>\
        <th><span class="label label-'+status+'">' + OrderStatus[gmOrder.orderStatus] + '</span></th>\
        </tr>'
    }
    $('#gmOrder tbody').html(html)
}



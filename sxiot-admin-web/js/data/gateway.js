$(function(){
    gateway_list()
    device_list()
})

function gateway_list() {
    var success = function (obj) {
        $('#gateway-list').empty()
        $.each(obj, function (i, val) {
            var list = `<tr>
                    <td>`+ val.gateway_id + `</td>
                    <td class="text-center">`+ val.gateway_name + `</td>
                    <td class="text-center"><span id="`+ val.gateway_id + `"></span></td>
                    <td class="text-center">
                    <button type="button" onclick="test_device('`+ val.gateway_id + `')" class="btn btn-sm btn-success">
                    <i class="icon-edit"></i> 模拟</button>
                        <button type="button" onclick="get_gateway('`+ val.gateway_id + `')" class="btn btn-sm btn-success">
                            <i class="icon-edit"></i> 编辑</button>
                        <button type="button" onclick="del_gateway('`+ val.gateway_id + `')"  class="btn btn-sm btn-danger">
                            <i class="icon-trash"></i> 删除</button>
                    </td>
                </tr>`
            $('#gateway-list').append(list)
            $('#' + val.gateway_id).qrcode({ width: 64, height: 64, text: val.gateway_id });
        });
    }
    AjaxReq(
        "get",
        "api/device/gateway/find",
        {},
        function () { },
        success,
        ReqErr
    )
}

$('#gateway-commit').click(function () {
    add_gateway()
});

$("#add-gateway").click(function () {
    clean_gateway()
})

function get_gateway(id) {
    AjaxReq(
        "get",
        "api/device/gateway/get",
        { gateway_id: id },
        function () { },
        function (msg) {
            $('#gateway-id').val(msg.gateway_id)
            $('#gateway-name').val(msg.gateway_name)
            $("#gateway-modal").modal("show")
        },
        ReqErr
    )
}

function del_gateway(id) {
    AjaxReq(
        "get",
        "api/device/gateway/delete",
        { gateway_id: id },
        function () { },
        function (msg) {
            gateway_list()
            ReqSuccess()
        },
        ReqErr
    )
}

function add_gateway() {
    var id = $('#gateway-id').val()
    var name = $('#gateway-name').val()

    var url = "api/device/gateway/create"
    if (id != "") {
        url = "api/device/gateway/update"
    }
    AjaxReq(
        "post",
        url,
        { gateway_name: name, gateway_id: id },
        function () { },
        function () {
            gateway_list()
            clean_gateway()
            $("#gateway-modal").modal("hide")
            ReqSuccess()
        },
        ReqErr
    )
}

function clean_gateway() {
    $('#gateway-id').val("")
    $('#gateway-name').val("")
}

var client
function test_device(id) {
    $("#test-content").empty()
    $("#test-modal").modal("show")
    DeviceType = new Array();
    client = new Paho.MQTT.Client(location.hostname, 8083, "device_"+id);
    // set callback handlers
    client.onConnectionLost = onConnectionLost;
    client.onMessageArrived = onMessageArrived;
    // connect the client
    client.connect({onSuccess:onConnect,userName:"6ba7b810-9dad-11d1-80b4-00c04fd430c8",password:"123456"});
}



// called when the client connects
function onConnect() {
  // Once a connection has been made, make a subscription and send a message.
  console.log("onConnect");
}

// called when the client loses its connection
function onConnectionLost(responseObject) {
  if (responseObject.errorCode !== 0) {
    console.log("onConnectionLost:"+responseObject.errorMessage);
  }
}

// called when a message arrives
function onMessageArrived(message) {
  console.log("onMessageArrived:"+message.payloadString);
  if (message.payloadString=="update"){
    var message = new Paho.MQTT.Message( $("#val-"+DeviceType[message.destinationName]).val());
    message.destinationName = $('#piont-'+DeviceType[message.destinationName]).val();
    client.send(message);
    return
  }
  $("#val-"+DeviceType[message.destinationName]).val(message.payloadString)
}

var device_points
function device_list() {
    var success = function (obj) {
        device_points=`<option value="">请选择一个数据点</option>`
        $.each(obj, function (i, val) {
            device_points = device_points+`<option value="` + val.device_id + `">` + val.device_alias + `</option>`
        });
    }
    AjaxReq(
        "get",
        "api/device/find",
        {},
        function () { },
        success,
        ReqErr
    )

}

var DeviceType = new Array();
$("#add-data").click(function () {
    var timestamp=new Date().getTime()
    var tr = `<tr>
                <td>
                    <select id="piont-`+timestamp+`" class="form-control">
                        `+device_points+`
                    </select>
                </td>
                <td>
                        <div class="input-group">
                                <input id="val-`+timestamp+`" type="text" value="0" class="form-control">
                                <span class="input-group-btn">
                                    <button id="up-`+timestamp+`" class="btn btn-default" type="button">更新</button>
                                </span>
                                <span class="input-group-btn">
                                    <button id="sui-`+timestamp+`" class="btn btn-default" type="button">随机</button>
                                </span>
                        </div>
                </td>
            </tr>`
    $("#test-content").append(tr)
    $('#up-'+timestamp).click(function(){
        var message = new Paho.MQTT.Message($('#val-'+timestamp).val());
        message.destinationName = $('#piont-'+timestamp).val();
        client.send(message);
    })
    $('#sui-'+timestamp).click(function(){
        var id =$('#piont-'+timestamp).val()
        setInterval('send_points("'+id+'")',5000); 
    })
    $('#piont-'+timestamp).change(function(){
        DeviceType[$('#piont-'+timestamp).val()]=timestamp
        client.subscribe($('#piont-'+timestamp).val());
    })
})

function send_points(id) {
    var message = new Paho.MQTT.Message((Math.random()*100).toFixed(2));
    message.destinationName = id
    client.send(message);
}

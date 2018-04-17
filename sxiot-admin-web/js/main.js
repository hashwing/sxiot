
var nodes 
$(function(){
    Getnodes()
    Getclients()
    Getcount()
})

function Getnodes(){
    AjaxReq(
        "get",
        "api/emq/cluster",
        {},
        function() {},
        function(msg) {
            nodes=msg
            $('#emq-num').html(msg.result.length+`<small>个</small>`)
        },
        ReqErr
    )
}

function Getclients(){
    AjaxReq(
        "get",
        "api/emq/client",
        {},
        function() {},
        function(msg) {
            var clients=0
            $.each(msg.result,function(i,val){
                var list= `<tr>
                <td>`+val.name+`</td>
                <td class="text-center">`+val.memory_used+`</td>
                <td class="text-center">`+nodes.result[i].uptime+`</td>
                <td class="text-center">`+val.clients+`</td>
                <td class="text-center">`+val.node_status+`</td>
            </tr>`
            $('#emq-list').append(list)
                clients+=val.clients
            })
           // $('#device-online').html(clients)
        },
        ReqErr
    )
}

function Getcount(){
    AjaxReq(
        "get",
        "api/emq/count",
        {},
        function() {},
        function(msg) {
            $("#device-sum").html(msg.gateway)
            $("#device-online").html(msg.gateway_online)
            $("#user-online").html(msg.user_online)
            $("#user-sum").html(msg.user)
            $("#data-sum").html(msg.device)
        },
        ReqErr
    )
}
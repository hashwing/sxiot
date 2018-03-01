$(function (){
    AjaxReq(
        "get",
        "api/user/get",
        {},
        function () { },
        function (msg) {
            $('#user-alias').html(msg.admin_alias)
        },
        ReqErr
    )
})
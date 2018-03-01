$(function () {
    user_list()
})

function user_list() {
    var success = function (obj) {
        $('#user-list').empty()
        $.each(obj, function (i, val) {
            var list = `<tr>
            <td>`+ val.admin_id + `</td>
                <td>`+ val.admin_alias + `</td>
                <td class="text-center">`+ val.admin_account + `</td>
                <td class="text-center">
                    <button type="button" onclick="get_user('`+ val.admin_id + `','` + val.admin_alias + `','` + val.admin_account + `')" class="btn btn-sm btn-success">
                        <i class="icon-edit"></i> 编辑</button>
                    <button type="button" onclick="del_user('`+ val.admin_id + `')" class="btn btn-sm btn-danger">
                        <i class="icon-trash"></i> 删除</button>
                </td>
            </tr>`
            $('#user-list').append(list)
        });
    }
    AjaxReq(
        "get",
        "api/user/find",
        {},
        function () { },
        success,
        ReqErr
    )
}

$('#user-commit').click(function () {
    add_user()
});

$('#user-add').click(function () {
    clean_user()
});
function add_user() {
    var id = $('#user-id').val()
    var name = $('#user-alias').val()
    var account = $('#user-account').val()
    var pwd = $('#user-password').val()
    AjaxReq(
        "post",
        "../api/user/create",
        { uid: id, name: name, account: account, pwd: pwd },
        function () { },
        function () {
            user_list()
            clean_user()
            $("#user-modal").modal("hide")
            ReqSuccess()
        },
        ReqErr
    )
}

function get_user(id, name, account) {
    $('#user-id').val(id)
    $('#user-alias').val(name)
    $('#user-account').val(account)
    $('#user-account').attr("disabled",true)
    $("#user-modal").modal("show")

}

function del_user(id) {
    AjaxReq(
        "get",
        "api/user/del",
        { uid: id },
        function () { },
        function () {
            user_list()
            ReqSuccess()
        },
        ReqErr
    )
}

function clean_user() {
    $('#user-id').val("")
    $('#user-alias').val("")
    $('#user-account').val("")
    $('#user-password').val("")
    $('#user-account').attr("disabled",false)
}


news_list()
function news_list(){
    var success =function(obj) {
        $('#message-list').empty()
        $.each(obj,function(i,val){
            var list= `<tr>
                <td>`+val.news_title+`</td>
                <td class="text-center">`+val.created+`</td>
                <td class="text-center">
                    <button type="button" onclick="del_news('`+val.news_id+`')" class="btn btn-sm btn-danger">
                        <i class="icon-trash"></i> 撤回</button>
                </td>
            </tr>`
            $('#message-list').append(list)
        });   
    }
    AjaxReq(
        "get",
        "api/news/find",
        {},
        function() {},
        success,
        ReqErr
    )
}

$('#message-commit').click(function(){
    add_news()
});

function add_news(){
    var title=$('#message-title').val()
    var content=$('#message-content').val()
    AjaxReq(
        "post",
        "api/news/create",
        {news_content:content,news_title:title},
        function() {},
        function() {
            news_list()
            clean_news()
            $("#message-modal").modal("hide")
            ReqSuccess()
        },
        ReqErr
    )
}


function del_news(id){
    AjaxReq(
        "get",
        "api/news/delete",
        {news_id:id},
        function() {},
        function() {
            news_list()
            ReqSuccess()
        },
        ReqErr
    )
}

function clean_news(){
    $('#message-id').val("")
    $('#message-title').val("")
    $('#message-content').val("")
}

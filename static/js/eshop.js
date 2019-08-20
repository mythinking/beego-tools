$(function () {
    $('#myTabs a').click(function (e) {
        e.preventDefault()
        if ($(this).prop('hash') != '#tmall') {
            $('#myModal-text').text('此功能暂未开放！！！');
            $('#myModal').modal('toggle');
            return
        }
        $(this).tab('show')
    });
    $("form").submit(function(){
        var storename = $('form #storename').val();
        var cookie = $('form #cookie').val();
        var title = $('form input[name="title"]:checked').length;
        var re = new RegExp("^[a-zA-Z0-9]+$");

        if (!storename || !re.test(storename)) {
            $('#myModal-text').text('请输入合法的店铺名，如uniqlo');
        } else if (!cookie || cookie.length < 10) {
            $('#myModal-text').text('请输入合法的cookie登录信息, 打开【天猫商城】登录后请求店铺，F12复制请求中[Request Headers][cookie]');
        } else if (title < 1) {
            $('#myModal-text').text('请至少选择一列');
        } else {
            createJob();
            return false;
        }

        $('#myModal').modal('toggle');
        return false;
    });
    
    function alertDialog(txt) {
        $('#myModal-text').text(txt);
        $('#myModal').modal('toggle');
    }

    function createJob() {
        var title = new Array();
        $('form input[name="title"]:checked').each(function(){
            title.push($(this).val());
        });
        $.post("/v1/crawlers/eshop",
        {
            storename:$('form #storename').val(),
            cookie:$('form #cookie').val(),
            title:title.join(','),
        },
        function(data,status){
            if (data) {
                if (data.Code !== 0) {
                    alertDialog(data.Msg);
                } 
            } 
            console.log(data, status)
        });
    }
    Line = 0;
    setInterval(function(){
        getLog(Line)
    }, 1000);
    function getLog(line) {
        line = line || 0;
        $.get("/v1/crawlers/log", {
            line:line
        },
        function(data, status){
            console.log(data);
            if (data) {
                Line = data.Data.line
                if (data.Data.content) {
                    $('#log').append(data.Data.content)
                    var ele = document.getElementById('log');
                    ele.scrollTop = ele.scrollHeight;
                }
            }
        });
    }
});
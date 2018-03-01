var data = {
    // labels 数据包含依次在X轴上显示的文本标签
    labels: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
    datasets: [{
        // 数据集名称，会在图例中显示
        label: "红队",

        // 颜色主题，可以是'#fff'、'rgb(255,0,0)'、'rgba(255,0,0,0.85)'、'red' 或 ZUI配色表中的颜色名称
        // 或者指定为 'random' 来使用一个随机的颜色主题
        color: "random",
        // 也可以不指定颜色主题，使用下面的值来分别应用颜色设置，这些值会覆盖color生成的主题颜色设置
        // fillColor: "rgba(220,220,220,0.2)",
        // strokeColor: "rgba(220,220,220,1)",
        // pointColor: "rgba(220,220,220,1)",
        // pointStrokeColor: "#fff",
        // pointHighlightFill: "#fff",
        // pointHighlightStroke: "rgba(220,220,220,1)",

        // 数据集
        data: [65, 59, 80, 81, 56, 55, 40, 44, 55, 70, 30, 40]
    }, {
        label: "绿队",
        color: "green",
        data: [28, 48, 40, 19, 86, 27, 90, 60, 30, 44, 50, 66]
    }]
};

var options = {}; // 图表配置项，可以留空来使用默认的配置

var myLineChart = $("#device-chart").lineChart(data, options);
var myLineChart = $("#user-chart").lineChart(data, options);
import logo from "~/assets/images/login_logofull.png"
var commaNumber = require("comma-number");
import moment from "moment";
moment.locale("th");

export default {
    async context(data, search) {
        var table = {
            style: "tableExample",
            margin: [10, 20, 10, 0],
            table: {
                widths: [30, 60, '*', 60, 60],
                body: [
                    [
                        { text: 'ลำดับ', style: 'tableHeader' },
                        { text: 'เลขแบบฟอร์ม', style: 'tableHeader' },
                        { text: 'รายการ', style: 'tableHeader' },
                        { text: 'ประเภท', style: 'tableHeader' },
                        { text: 'จำนวนเงิน', style: 'tableHeader' },
                    ],
                ]
            },
        }
        data.map((d, i) => {
            table.table.body.push([
                {
                    text: i + 1,
                    alignment: "center"
                },
                {
                    text: d.number_form,
                    alignment: "center"
                },
                {
                    text: d.title,
                },
                {
                    text: d.type.nametype,
                    alignment: "center"
                },
                {
                    text: commaNumber(d.total),
                },
            ])
        });
        var content = [
            {
                image: await fnUrlToB64(logo),
                width: 200,
                alignment: "center",
                margin: [0, 0, 0, 0],
            },
            [
                {
                    text: `รายงานสรุปรายการ ${search.type != null ? data[0].type.nametype : 'รายรับ รายจ่าย เบิก ยืม คืน'} `,
                    style: 'header',
                    margin: [0, 20, 0, 0],
                },
                {
                    text: `แผนกวิชาเทคโนโลยีสารสนเทศ วิทยาลัยเทคนิคเชียงใหม่`,
                    style: 'header',
                },
                {
                    text: `ประจำเดือน ${moment(search.month).format('MMMM')} ปี ${moment(search.year).add(543, "year").format('YYYY')}`,
                    style: 'header',
                }],
            table
        ];
        var docDefinition = {
            content: content,
            styles: {
                header: {
                    fontSize: 12,
                    bold: true,
                    alignment: 'center',
                    margin: [0, 2, 0, 0]
                },
                tableHeader: {
                    bold: true,
                    fontSize: 9,
                    color: "black",
                    alignment: 'center'
                },
                tableExample: {
                    fontSize: 8,
                    color: "black"
                },
                tableMargin: {
                    margin: [0, 0, 0, 0]
                }
            },
            defaultStyle: {
                font: "Sarabun",
                fontSize: 8
            },
            pageOrientation: "portrait",
            pageMargins: [30, 50, 30, 80],
            footer: function (currentPage, pageCount) {
                if (currentPage == pageCount)
                    return {
                        margin: [10, 0, 10, 0],
                        height: 80,
                        columns: [
                            [
                                {
                                    columns: [
                                        [
                                            {
                                                columns: [
                                                    {
                                                        text: "ลงชื่อ",
                                                        alignment: "right",
                                                        width: 50
                                                    },
                                                    {
                                                        text:
                                                            "____________________________________________________"
                                                    },
                                                    {
                                                        text: "เจ้าหน้าที่การเงิน"
                                                    }
                                                ]
                                            },
                                            {
                                                columns: [
                                                    {
                                                        text: "",
                                                        width: 50
                                                    },
                                                    {
                                                        text:
                                                            "(____________________________________________________)"
                                                    }
                                                ]
                                            }
                                        ],
                                        [
                                            {
                                                columns: [
                                                    {
                                                        text: "ลงชื่อ",
                                                        alignment: "right",
                                                        width: 50
                                                    },
                                                    {
                                                        text:
                                                            "____________________________________________________"
                                                    },
                                                    {
                                                        text: "หัวหน้าแผนก"
                                                    }
                                                ]
                                            },
                                            {
                                                columns: [
                                                    {
                                                        text: "",
                                                        width: 50
                                                    },
                                                    {
                                                        text:
                                                            "(____________________________________________________)"
                                                    }
                                                ]
                                            }
                                        ]
                                    ]
                                },
                            ]
                        ]
                    };
            }
        };

        return docDefinition;
    }
}

function fnUrlToB64(event) {
    var httpRequest = new XMLHttpRequest();
    return new Promise((resolve, reject) => {
        httpRequest.onload = function () {
            var reader = new FileReader();
            var Bind_HttpReq = this;
            reader.onload = event => {
                if (Bind_HttpReq.status == 200) {
                    resolve(event.target.result);
                } else {
                    console.log(
                        `200 Error XMLHttpRequest ${Bind_HttpReq.statusText}`
                    );
                }
            };
            reader.onerror = () => {
                console.log(
                    `On Error XMLHttpRequest ${Bind_HttpReq.statusText} >> ${reject}`
                );
            };
            reader.readAsDataURL(httpRequest.response);
        };
        httpRequest.open("GET", event);
        httpRequest.responseType = "blob";
        httpRequest.send();
    });
}
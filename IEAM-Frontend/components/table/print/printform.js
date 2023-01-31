import logo from "~/assets/images/login_logofull.png"
var commaNumber = require("comma-number");
import moment from "moment";
moment.locale("th");

export default {
    async context(data) {
        console.log("Data ------>", data);
        // console.log(data.map((d) => {
        //     return {
        //         text: d,
        //         style: 'tableHeader'
        //     }
        // }));
        var table = {
            style: "tableExample",
            margin: [0, 10, 0, 0],
            table: {
                widths: [30, '*', 30, 30, 50, 60],
                body: [
                    [
                        { text: 'ลำดับ', style: 'tableHeader' },
                        { text: 'รายการ', style: 'tableHeader' },
                        { text: 'จำนวน', style: 'tableHeader' },
                        { text: 'หน่วย', style: 'tableHeader' },
                        { text: 'ราคา/หน่วย', style: 'tableHeader' },
                        { text: 'รวม', style: 'tableHeader' },
                    ],
                    // [{ text: 'total', style: 'tableHeader', colSpan: 4, alignment: 'center' }, {}, {}, {}, { text: '10,000', alignment: 'center' }],
                ]
            },
        }
        data.list.map((d, i) => {
            table.table.body.push([
                {
                    text: i + 1,
                    alignment: "center"
                },
                {
                    text: d.title,

                },
                {
                    text: commaNumber(d.number),
                    alignment: "center"
                },
                {
                    text: d.unit,
                    alignment: "center"
                },
                {
                    text: commaNumber(d.unit_price),
                    alignment: "center"
                },
                {
                    text: commaNumber(d.total),
                },
            ])
        });
        table.table.body.push([{ text: 'ยอดรวมทั้งหมด', style: 'tableHeader', colSpan: 5, alignment: 'center' }, {}, {}, {}, {}, { text: commaNumber(data.total), alignment: 'center' }])
        var content = [
            {
                image: await fnUrlToB64(logo),
                width: 150,
                alignment: "center",
                margin: [0, 0, 0, 30],
            },
            [
                {
                    text: `ใบ${data.type.nametype}เงิน`,
                    style: 'header',
                    margin: [0, 0, 0, 0],
                },
                {
                    text: `แผนกวิชาเทคโนโลยีสารสนเทศ วิทยาลัยเทคนิคเชียงใหม่`,
                    style: 'header',
                },
            ],
            {
                columns: [
                    [{
                        text: `หัวเรื่อง : ${data.title}`,
                        style: 'title',
                    },
                    {
                        text: `เลขแบบฟอร์ม : ${data.number_form}`,
                        style: 'title',
                    }],
                    [{
                        text: `วันที่ออก : ${moment(data.CreatedAt).format('DD/MMMM/YYYY')}`,
                        style: 'title',
                    },
                    {
                        text: `จำนวนเงิน : ${commaNumber(data.total)}`,
                        style: 'title',
                    }],
                ],
                margin: [0, 20, 0, 20]
            },
            {
                text: `ชื่อผู้ยืนเอกสาร : ${data.user.firstname + " " + data.user.lastname }`,
                style: 'title',
            },
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
                title: {
                    bold: true,
                    fontSize: 9,
                },
                tableExample: {
                    fontSize: 8,
                    color: "black"
                },
                tableHeader: {
                    bold: true,
                    alignment: "center"
                }
            },
            defaultStyle: {
                font: "Sarabun",
                fontSize: 8
            },
            pageOrientation: "portrait",
            pageMargins: [50, 50, 50, 80],
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
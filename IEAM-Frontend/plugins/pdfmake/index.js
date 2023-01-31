import Vue from "vue";
import pdfMake from "pdfmake";
import pdfFonts from "assets/fonts.js";

pdfMake.vfs = pdfFonts.pdfMake.vfs;
pdfMake.fonts = {
    Sarabun: {
        normal: "Sarabun-Regular.ttf",
        bold: "Sarabun-Medium.ttf",
        italics: "Sarabun-Italic.ttf",
        bolditalics: "Sarabun-MediumItalic.ttf"
    }
};
Vue.use(pdfMake);

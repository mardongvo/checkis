# -*- coding: utf-8 -*-

import zipfile
import xml.dom as dom
import xml.dom.minidom as minidom
import shutil
import os

def iterN(inp_list, n):
    if len(inp_list)<n: return
    for i in range(len(inp_list)-n+1):
        yield inp_list[i:i+n]

def getT(node):
    for e in node.getElementsByTagName("w:t"):
        return e.firstChild

def joinT(parent, nX):
    getT(nX[0]).data = "".join(map(lambda n: getT(n).data, nX))
    for node in nX[1:]:
        parent.removeChild(node)

def checkX(nX):
    if (getT(nX[0]).data.endswith("#") and getT(nX[-1]).data.startswith("#")) or \
       (getT(nX[0]).data.startswith("#") and getT(nX[-1]).data.endswith("#")):
        pass
    else:
        return False
    for i in range(1,len(nX)-1):
        t = getT(nX[i]).data
        if t.find(" ") != -1: return False
        if t.find("#") != -1: return False
    return True

def docfix(doc):
    for tag in ["w:proofErr", "w:bookmarkStart", "w:bookmarkEnd"]:
        for e in doc.getElementsByTagName(tag):
            p = e.parentNode
            if p != None:
                p.removeChild(e)
    for par in doc.getElementsByTagName("w:p"):
        for n in range(2,8):
            doit = True
            while doit:
                doit = False
                for nX in iterN(par.getElementsByTagName("w:r"), n):
                    if checkX(nX):
                        doit = True
                        joinT(par,nX)
                        break
    return doc

def fixfile(filepath):
    #backup
    shutil.copy(filepath, filepath+".bak")
    zf = zipfile.ZipFile(filepath, 'r')
    zf_out = zipfile.ZipFile("temp.docx", 'w', zipfile.ZIP_DEFLATED)
    for i in zf.namelist():
        inp = zf.open(i, 'r')
        out = zf_out.open(i, 'w')
        if i == "word/document.xml":
            doc = minidom.parse(inp)
            docfix(doc)
            out.write(doc.toxml().encode("UTF-8"))
        else:
            out.write(inp.read())
        inp.close()
        out.close()
    zf.close()
    zf_out.close()
    shutil.copy("temp.docx", filepath)
    os.remove("temp.docx")

fixfile("template_docs_act.docx")
fixfile("template_docs_req.docx")
fixfile("template_docs_decision.docx")
fixfile("template_docs_charge.docx")

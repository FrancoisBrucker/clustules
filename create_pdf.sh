#! /bin/sh

dot -Tpdf -O treillis.dot
dot -Tpdf -O non_etirable_treillis.dot
neato -Tpdf -O non_etirable.dot
neato -Tpdf -O non_etirable_label.dot
neato -Tpdf -O non_etirable_nuf.dot
neato -Tpdf -O non_etirable_nuf_label.dot

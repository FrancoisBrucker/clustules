#! /bin/sh

cd out

dot -Tpdf -O   orig_intervalles_treillis.dot
dot -Tpdf -O   orig_non_etirable_treillis.dot
neato -Tpdf -O orig_g.dot

dot -Tpdf -O   nu_intervalles_treillis.dot
dot -Tpdf -O   nu_non_etirable_treillis.dot
neato -Tpdf -O nu_g.dot

cd ..
#!/bin/bash

acbuild begin # Iniciando o acbuild
acbuild set-name perylemke/hello-rkt # Determina um nome a nossa imagem
acbuild copy hello-rkt /bin/hello-rkt # Copiando o binário para o acbuild
acbuild set-exec /bin/hello-rkt # Deixando como executável
acbuild port add www tcp 5000 # Expõe a porta 5000
acbuild label add version 0.0.1 # Determina uma versão
acbuild label add arch amd64 # Em qual arquitetura ela roda
acbuild label add os linux # Em qual sistema operacional roda
acbuild annotation add authors "Pery Lemke <pery.lemke@gmail.com>" # Determina o Autor
acbuild write hello-rkt-0.0.1-linux-amd64.aci # Dá um nome ao tar.gz da imagem
acbuild end # Finaliza o acbuild.

exit 0

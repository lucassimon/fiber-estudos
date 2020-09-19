#!/bin/bash

diretorioAplicacao="/usr/local/gocep/"
nomeArquivoAplicacao="go-cep"
nomeArquivoConfiguracao=".env"

nomeServico="go-cep"
nomeArquivoServico="go-cep.service"
caminhoScriptServicoLinux="/etc/systemd/system/"

echo "GO-CEP - Serviço de consulta de endereço por CEP"
echo 

if [ -e $diretorioAplicacao ] && rm -R $diretorioAplicacao; then
    echo "[√] Removido instalação antiga"
fi

if [ -e "$caminhoScriptServicoLinux$nomeArquivoServico" ] && rm -R "$caminhoScriptServicoLinux$nomeArquivoServico"; then
    echo "[√] Removido serviço antigo"
fi

if cp -R ./ "$diretorioAplicacao"; then 
    echo "[√] Arquivos da aplicação copiado para $diretorioAplicacao"
else
    echo
    echo "[!] Erro ao copiar oo arquivos da aplicação para $diretorioAplicacao"
    echo

    exit 1
fi

if chmod +x $diretorioAplicacao$nomeArquivoAplicacao; then 
    echo "[√] Permissões setadas"
else 
    echo
    echo "[!] Erro ao dar permissão de execução para $diretorioAplicacao$nomeArquivoAplicacao"
    echo

    exit 1
fi 

if cp $nomeArquivoServico $caminhoScriptServicoLinux$nomeArquivoServico; then 
    echo "[√] Arquivo $nomeArquivoServico copiado para $caminhoScriptServicoLinux$nomeArquivoServico"
else 
    echo
    echo "[!] Erro ao copiar arquivo $nomeArquivoServico copiado para $caminhoScriptServicoLinux$nomeArquivoServico"
    echo

    exit 1
fi

if systemctl daemon-reload; then 
    echo "[√] Atualizando lista de serviços do sistema operacional"
else 
    echo
    echo "[!] Erro ao atualizar lista de serviços do sistema operacional"
    echo

    exit 1
fi

if service $nomeServico start; then
    echo "[√] Executando serviço"
else
    echo
    echo "[!] Falha ao iniciar o serviço, execute [journalctl -u $nomeServico -f] para mais detalhes"
    echo

    exit 1
fi

echo
echo
service $nomeServico status

echo 
echo "Fim da instalação, em caso de problema execute o comando: [ journalctl -u $nomeServico -f ] para mais detalhes"

#!/usr/bin/env bash
# Falha se algum serviço depender de outro fora do grafo permitido,
# ou se importar internal/ de outro serviço. Compatível com bash 3.2 (macOS).
set -euo pipefail
cd "$(dirname "$0")/../server"
BASE=github.com/igorrafae/MIGDOL/server/services

allowed() { # serviço -> dependências permitidas
  case "$1" in
    audit)     echo "" ;;
    directory) echo "audit" ;;
    delivery)  echo "directory audit" ;;
    blobs)     echo "directory" ;;
    gateway)   echo "directory delivery" ;;
    *)         echo "__desconhecido__" ;;
  esac
}

fail=0
for dir in services/*/; do
  svc=$(basename "$dir")
  ok=" $(allowed "$svc") "
  if [ "$ok" = " __desconhecido__ " ]; then echo "PROIBIDO: serviço $svc não está no grafo"; fail=1; continue; fi
  for d in $(grep -oE "$BASE/[a-z]+ " "$dir/go.mod" | sed "s#$BASE/##; s# ##" || true); do
    case "$ok" in *" $d "*) ;; *) echo "PROIBIDO: $svc depende de $d"; fail=1 ;; esac
  done
  if grep -rn "services/[a-z]*/internal" "$dir" --include=*.go | grep -v "services/$svc/internal"; then
    echo "PROIBIDO: $svc importa internal/ de outro serviço"; fail=1
  fi
done
[ "$fail" -eq 0 ] && echo "fronteiras ok"
exit "$fail"

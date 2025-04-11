# Api EKS

## Visão Geral
Esta API foi desenvolvida para gerenciar recursos Kubernetes, permitindo realizar operações de **scale up** em **Deployments** e **StatefulSets**.

Ela funciona em conjunto com o [kube-downscaler](https://codeberg.org/hjacobs/kube-downscaler), automatizando o processo de zerar os recursos fora do horário comercial e permitindo que usuários escalem manualmente os recursos quando necessário usando essa API desenvolvida.

Desta forma conseguimos diminuir custos com os nodes.

## Pré-requisitos
- Lista de pré-requisitos necessários para usar a API (ex: Node.js, Python, etc.)
- Como instalar dependências

## Instalação
Passos para clonar o repositório e instalar as dependências.

```bash
git clone https://github.com/jamessantos10/api-sts-eks.git
cd api-sts-eks
# Comando para instalar dependências, por exemplo:
npm install

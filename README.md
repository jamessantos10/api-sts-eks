# Api EKS

## Visão Geral
Esta API foi desenvolvida para gerenciar recursos Kubernetes, permitindo realizar operações de **scale up** em **Deployments** e **StatefulSets**.

Ela funciona em conjunto com o [kube-downscaler](https://codeberg.org/hjacobs/kube-downscaler), automatizando o processo de zerar os recursos fora do horário comercial e permitindo que usuários escalem manualmente os recursos quando necessário usando essa API desenvolvida.

Desta forma conseguimos diminuir custos com os nodes.

## Pré-requisitos
- Kubernetes configurado e acessível.
- Ferramentas necessárias instaladas:
  - `kubectl`
  - Configuração de acesso ao cluster Kubernetes.
- Node.js.

## Instalação
Clone o repositório e instale as dependências necessárias.

```bash
git clone https://github.com/jamessantos10/api-sts-eks.git
cd api-sts-eks

Certifique-se de ter o Go instalado em sua máquina. Em seguida, execute:

go build
./api-sts-eks

A API estará disponível na porta 8080.




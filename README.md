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
```

## Uso

Para utilizar a API, envie um POST para a rota /up com o JSON especificado:

```
curl -X POST http://localhost:8080/up -H "Content-Type: application/json" -d '{"power": true, "idaccount": "123456789012"}'
```

Essa rota é responsável por executar um kubectl no deployment do kubedownscaler alterando o horário de disponibilidade do Cluster e escalando novamente os deployments e statefullsets.

# Observação

Essa aplicação está usando AssumeRole, para testar 100% de forma local, é necessário usar credencias locais para testes. Por padrão, o método AssumeRole é utilizado para criar credenciais com a conta de destino. Para usar credenciais locais, substitua o trecho de código abaixo [client]([https://codeberg.org/hjacobs/kube-downscaler](https://github.com/jamessantos10/api-sts-eks/blob/main/api/src/controllers/sts.go#L26)):

```
// Criar credenciais com a conta de destino, usando o método AssumeRole
creds := stscreds.NewCredentials(sess, "", func(p *stscreds.AssumeRoleProvider) {
    p.RoleARN = *aws.String(role)
    p.RoleSessionName = *aws.String("600")
})
```

Alteração para Credenciais Locais:

```
// Usar credenciais locais para testes
creds := sess.Config.Credentials
```


# Detalhes Técnicos

## Conexão STS
A API utiliza o AWS STS (Security Token Service) para criar uma conexão temporária com a conta de destino, baseada no idaccount fornecido. Certifique-se de que as permissões necessárias estão configuradas para permitir a operação.

## Geração do .kubeconfig
Após estabelecer a conexão STS, a API cria um client para o cluster Kubernetes associado à conta de destino. O arquivo .kubeconfig é gerado para permitir a execução de comandos kubectl.

## Execução do kubectl
Com o .kubeconfig gerado, a API pode executar comandos kubectl para interagir com o cluster e ligar os serviços.

Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

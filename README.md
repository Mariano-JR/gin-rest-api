# API REST criada com GIN

## API criada durante o curso de Go e Gin: Criando API Rest com simplicidade oferecido pela Alura, onde posteriormente também foi utilizado como base para o curso Go: Validações, teste e páginas HTML.
## Se trata de uma API de armazenagem de dados de alunos para uma escola, onde a Struct tem os parametros de ID, Nome, CPF e RG.
## Nestes cursos foram abordados os conceitos do framework Go, Gin. Também foram utilizadas bibliotecas externas como o ORM Gorm, Testify e Validator V2 para testes e validações.

### Através do Docker, foi criado um banco de dados Postgres para a armazenagem dos dados dos alunos onde é possivel fazer as seguintes requisições:

> Criar um aluno através da rota "/alunos" com o método POST;

> Editar um aluno através da rota "/alunos/id-do-aluno" com o método PATCH;

> Excluir um aluno através da rota "/alunos/id-do-aluno" com o método DELETE;

> Buscar todos os alunos através da rota "/alunos" com o método GET;

> Buscar um aluno por ID através da rota "/alunos/id-do-aluno" com o método GET;

> Buscar um aluno por CPF através da rota "/alunos/cpf-do-aluno" com o método GET;

> Acessa a página HTML pela rota "/index" com o método GET.

### Foi desenvolvido um ambiente de testes para essa API, onde foi possivel fazer verificações de dados passados ao banco de dados, onde o parametro Nome não pode ser em branco, CPF deve ter 11 caracteres dentre 0-9 e o RG deve ter 9 caracteres tambem dentre 0-9.
### Nesses testes são criados alunos teste que após a verificação, é deletado automaticamente do banco de dados.
### Nesses testes é possivel verificar se o StatusCode é o esperado para a requisição, Se o Nome, CPF e RG passados como parametro batem com o do aluno no banco de dados.
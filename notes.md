Como funciona?

Para poder extrair alguma info do json da response, é necessário criar uma struct com os campos aos quais serão interessantes a serem usados ou processados posteriormente. A mesma coisa é feita quando se trata um post request


a função Decode(), do pacote "encoding/json", faz essa conversão do json encoded, para o objeto json padrao
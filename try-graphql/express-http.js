//https://www.bilibili.com/video/BV154411o7TD?from=search&seid=17169381987330463339
//https://github.com/sunwu51/graphql-demo
//https://graphql.org/code/
var express = require('express');
var { graphqlHTTP } = require('express-graphql');
var { buildSchema } = require('graphql');
 
var schema = buildSchema(`
  type Query {
    hello: String
  }
`);
 
var root = { hello: () => 'Hello world!' };
 
var app = express();
app.use('/graphql', graphqlHTTP({
  schema: schema,
  rootValue: root,
  graphiql: true,
}));
app.listen(4000, () => console.log('Now browse to localhost:4000/graphql'));
/* 在GraphiQL框中输入如下执行可得结果
query{
	hello
}
*/

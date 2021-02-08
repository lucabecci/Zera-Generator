> Read the all file for the correct usage.

### How to run the server:

Configure the package.json with your scripts or use `node src/index.js` (if you use a javscript template) or `tsc watch src/index.ts` (if you use a typescript template)

### Port of the server:

The default port is 4000, you can change it if you want.

### Configure the database:

- If you use Javascript/Typescript and mongoose you need to:

```jsx
await mongoose.connect(
'mongodb://localhost:27017/yourapp', //insert your db url connection
{useNewUrlParser: true} // insert more properties if you need 
)
```

- If you use Typescript and typeORM you need to:

```tsx
await createConnection({
    type: "", //your type of db
    host: "localhost", //the host of the db
    port: 5672, //port of your db
    username: "postgres", //db username
    password: "postgres", //psw username
    database: "yourapp", //name of your db
    synchronize: true, //i not recommend use this in production
    logging: true, //for see the logs of your db
    entities: [], //on here you can insert your entities
})
```

- The Graphql template will be included in the next release.
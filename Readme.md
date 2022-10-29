# FullCycle Graphql Go Application

## Technologies

- Go 1.19
- Sqlite3
- [Gqlgen](https://gqlgen.com/getting-started/)

## Tips

If you want to generate new models and resolver based on your schema, just run this command of root folder of projet
`go run github.com/99designs/gqlgen generate`

If you had some issues using sqlite3 lib to connect to your database on windows you had to install this program in your computer
[tdmgcc](https://github.com/jmeubank/tdm-gcc/releases/download/v10.3.0-tdm64-2/tdm64-gcc-10.3.0-2.exe)

## Mutations and queries

```graphql
mutation createCategory {
  createCategory(input: {
    name: "Technology", 
    description: "technology courses"
  }) {
    id
    name
    description
  }
}

mutation createCourse {
  createCourse(input: {
    name: "Go lang", 
    description: "Go lang course",
    categoryId: "4987d819-9914-4396-a30c-3c7b3966456e"
  }) {
    id
    name
    description
  }
}


query QueryCategories {
    categories {
    id
    name
    description
	}
}

query QueryCategoriesWithCourses {
    categories {
      id
      name
      description,
      courses {
        id
        name
        description
      }
	}
}

query QueryCourses {
    courses {
    id
    name
	}
}

```

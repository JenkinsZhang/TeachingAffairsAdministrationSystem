.
├── backend
│   ├── api.md
│   ├── pkg
│   │   └── windows_amd64
│   │       ├── func
│   │       │   ├── admin.a
│   │       │   ├── login.a
│   │       │   ├── student.a
│   │       │   └── teacher.a
│   │       ├── models.a
│   │       ├── router.a
│   │       └── utils.a
│   ├── school.md
│   ├── school.sql
│   └── src
│       ├── Dockerfile
│       ├── .dockerignore
│       ├── func
│       │   ├── admin
│       │   │   ├── courseManagement.go
│       │   │   ├── profile.go
│       │   │   ├── studentManagement.go
│       │   │   ├── teacherManagement.go
│       │   │   └── termManagement.go
│       │   ├── login
│       │   │   └── login.go
│       │   ├── student
│       │   │   ├── courseCalendar.go
│       │   │   ├── courseQuery.go
│       │   │   ├── profile.go
│       │   │   ├── scoreCollege.go
│       │   │   ├── scoreSummary.go
│       │   │   ├── scoreTable.go
│       │   │   └── scoreTrend.go
│       │   ├── teacher
│       │   │   ├── classTable.go
│       │   │   ├── courseCalendar.go
│       │   │   ├── profile.go
│       │   │   ├── scoreAnalysis.go
│       │   │   └── scoreManagement.go
│       │   └── tools
│       │       └── tools.go
│       ├── .gitignore
│       ├── go.mod
│       ├── go.sum
│       ├── main.go
│       ├── models
│       │   └── models.go
│       ├── router
│       │   └── router.go
│       ├── test
│       │   └── test.go
│       └── utils
│           ├── check.go
│           ├── connect.go
│           ├── delete.go
│           ├── insert.go
│           ├── parse.go
│           ├── query.go
│           ├── queryStudent.go
│           ├── queryTeacher.go
│           ├── score.go
│           ├── term.go
│           ├── Token.go
│           ├── tools.go
│           └── update.go
├── docker-compose.yml
├── frontend
│   ├── assets
│   │   ├── img
│   │   │   └── pexels-photo-210598.jpeg
│   │   ├── js
│   │   │   ├── calendarColumnsMock.js
│   │   │   ├── scoreManagementMock.js
│   │   │   └── tokenTools.js
│   │   ├── json
│   │   │   ├── classColumns.json
│   │   │   ├── identity.json
│   │   │   └── menu.json
│   │   ├── README.md
│   │   └── style
│   │       ├── main.scss
│   │       └── reset.scss
│   ├── .babelrc
│   ├── components
│   │   ├── Logo.vue
│   │   ├── README.md
│   │   └── sideMenu
│   │       ├── customMenu
│   │       │   ├── assist.js
│   │       │   ├── customMenu.vue
│   │       │   ├── emitter.js
│   │       │   ├── link.js
│   │       │   ├── menu-item.vue
│   │       │   ├── mixin.js
│   │       │   └── submenu.vue
│   │       ├── sideMenu.scss
│   │       └── sideMenu.vue
│   ├── Dockerfile
│   ├── .dockerignore
│   ├── .editorconfig
│   ├── .gitignore
│   ├── layouts
│   │   ├── default.vue
│   │   └── README.md
│   ├── middleware
│   │   ├── anonymous.js
│   │   ├── authenticated.js
│   │   └── checkAuth.js
│   ├── nuxt.config.js
│   ├── package.json
│   ├── pages
│   │   ├── index.vue
│   │   ├── login.scss
│   │   ├── login.vue
│   │   ├── _users
│   │   │   ├── classTable.vue
│   │   │   ├── courseAnalysis.vue
│   │   │   ├── courseCalendar.vue
│   │   │   ├── courseManagement.vue
│   │   │   ├── courseQuery.vue
│   │   │   ├── index.vue
│   │   │   ├── profile.vue
│   │   │   ├── scoreAnalysis.vue
│   │   │   ├── scoreCollege.vue
│   │   │   ├── scoreManagement.vue
│   │   │   ├── scoreSummary.vue
│   │   │   ├── scoreTable.vue
│   │   │   ├── scoreTrend.vue
│   │   │   ├── studentManagement.vue
│   │   │   ├── teacherManagement.vue
│   │   │   └── termManagement.vue
│   │   ├── users.scss
│   │   └── _users.vue
│   ├── plugins
│   │   ├── axios.js
│   │   ├── cookies.client.js
│   │   ├── dayjs.js
│   │   ├── iview.js
│   │   └── vcharts.client.js
│   ├── .prettierrc
│   ├── README.md
│   ├── server
│   │   └── index.js
│   ├── static
│   │   ├── favicon.ico
│   │   ├── icon.png
│   │   ├── README.md
│   │   └── sw.js
│   ├── store
│   │   └── index.js
│   └── yarn.lock
├── .gitignore
└── README.md

32 directories, 122 files

export default `
schema {
  query: Query
  mutation: Mutation
}

type Query {
  student(username: String!, password: String!): Student
  authenticatedStudent(token: String!): Student
}

type Mutation {
  login(username: String!, password: String!): AuthenticationData
}

type Student {
  isAuthorized: Boolean
  courses(code: String): [Course]
  schedule: [Slot]
  transcript: Transcript
}

type AuthenticationData {
  isAuthorized: Boolean
  token: String
}

type Course {
  code: String
  name: String
  absence: Absence
  coursework: [Component]
  midterm: Midterm
  exam: Exam
}

type Absence {
  level: Int
  severity: AbsenceSeverity
}

type Slot {
  course: Course
  type: SlotType
  weekday: SlotWeekday
  number: Int
  venue: Venue
}

type Venue {
  room: String
  building: String
}

type Component {
  type: String
  grade: Float
  maximumGrade: Float
}

type Midterm {
  grade: Float
}

type Exam {
  venue: Venue
  seat: String
  startsAt: String
}

type Transcript {
  cumulativeGPA: Float
  semesters: [Semester]
}

type Semester {
  year: Int
  type: SemesterType
  gpa: Float
  entries: [Entry]
}

type Entry {
  course: Course
  grade: Grade
  creditHours: Int
}

type Grade {
  german: Float
  american: String
}

enum SlotWeekday {
  SATURDAY
  SUNDAY
  MONDAY
  TUESDAY
  WEDNESDAY
  THURSDAY
}

enum SlotType {
  LECTURE
  TUTORIAL
  LAB
}

enum AbsenceSeverity {
  HIGH
  MEDIUM
  LOW
}

enum SemesterType {
  WINTER
  SPRING
  SUMMER
}
`;

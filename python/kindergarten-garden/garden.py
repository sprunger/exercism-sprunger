class Garden:
  def __init__(self, diagram, students = None):
    if students:
      self.students = sorted(students)
    else:
      self.students = ["Alice", "Bob", "Charlie", "David", "Eve",
          "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid",
          "Larry"]

    self.plant_names = {'C': 'Clover', 'G': 'Grass', 'R': 'Radishes', 'V': 'Violets'}

    self.rows = diagram.split()

  def plants(self, student):
    student_start = self.students.index(student) * 2
    student_end = student_start + 2

    student_plants = [self.plant_names[letter]
      for letter in (
        self.rows[0][student_start:student_end] +
        self.rows[1][student_start:student_end]
      )]

    return student_plants

# Below is an implementation by @mariazverina
# I spent several hours trying to figure out how to get a dict with
# student name and their pots.  This is brilliant.
#
# class Garden:
#     STUDENTS = ['Alice', 'Bob', 'Charlie', 'David', 'Eve', 'Fred',
#         'Ginny', 'Harriet', 'Ileana', 'Joseph', 'Kincaid', 'Larry']

#     FLOWERS = {'C': 'Clover', 'R': 'Radishes',
#         'G': 'Grass', 'V': 'Violets'}

#     def __init__(self, garden, students=STUDENTS):
#         garden = garden.split()
#         flowers = zip(*[iter(garden[0])]*2 + [iter(garden[1])]*2)
#         self.pots = dict(zip(sorted(students), flowers))

#     def plants(self, who):
#         return map(self.FLOWERS.get, self.pots[who])

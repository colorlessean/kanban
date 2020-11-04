# Kanban Main Design Document and Ideas Board

## Main Object Structure
Inteded to have the concepts of boards, columns and tickets at the for-front.

Ticket objects need some sub objects and attributes that are required for their general function. Tickets will need the concept of tags to order by categories in a visual and functional way. This tags will exist as concepts seperate from tickets themselves.

Tickets also need due dates, comments, sub-task lists.

Tickets will belong to a column or list that is present on the board. They can also exist in the special 'archive' list which will not show them as displayed but they can be restored from as opposed to fully deleted from.

Boards will contain tickets, and columns. They will also have the concept of a title and maybe an icon that is all. Could also have a category of their own perhaps to allow top level sorting by general intention. I.e. work vs personal vs. project1 vs. project 2, etc. We will see as time goes on.

## Minimal Viable Product
Simply a list of boards of which tickets are contained into columns. Tickets will have a title, body and bullet list of tasks which can be added or removed. They will be persistently stored in a database for access later. Many boards can be available for usage.

Good place to start is going to be a single Board layout to drag and drop around a set of tickets. Will then expand to the MVP after that.

## Product Stack
Want to use a Go back-end to handle the MVC and the API strucutre.

Hopefully can make use of nginx when hosting comes into play on a rasperry pi or a VPS. May be sufficient to just make use of Gorilla

Want to use a react front end or Angular front end to handle the visuals of the tickets in lists. I really want to focus on having a minimal front-end with an easy to use interface that is mobile compatible hence the usage of a library. CSS wise I see no reason to go anything more complex than twitter bootstrap.

See no reason for anything more complex than a collection of SQLite tables.

## Database Schema
The database schema should reasonable follow the object structure to store all of the intended values and concepts.

### Tables:

Database Tables with their key layouts.

- Boards
    - BoardID (Primary Key)
    - TagID (Foreign key)
    - title
- BoardTags
    - BoardTagID (Primary Key)
    - tagName
    - color

- Columns
    - ColumnID (Primary Key)
    - BoardID (Foreign Key)
    - columnName

- Tickets
    - TicketID (Primary Key)
    - ColumnID (Foreign Key)
    - TicketTagID
    - Title
    - BodyText
- TicketTags
    - TicketTagID
    - tagName
    - color
- SubTask
    - SubTaskID (Primary Key)
    - TicketID (Foreign Key)
    - completeBool
    - order
    - taskBody
- Notes
    - NoteID (Primary Key)
    - TicketID (Foreign Key)
    - timeStamp/order
    - noteBody

## Back-End Architecture

Main Idea is that each table will have a model view and controller.

### API endpoints

/api/board/{boardID}
- Create a new board
- Update a boards name or tag
- Delete a board and all its subtickets
- Get a list of boards
- Get a list of boards, column, and tickets

/api/board/tags/{boardTagID}
- Create a new board tag
- Update a tag name / color
- Delete a tag and clear the tags from all boards that use it
- Get a list of tags

/api/board/column/{columnID}
- Create a new column
- Update a column name
- Delete a column and all the tickets in the column (ensure user is notified)

/api/ticket/{ticketID}
- Create a new ticket with title, collaborators, etc.
- Update a ticket with any form of information
- Get a ticket by its id
- Delete a ticket

### Webforms

1. Ticket Creation Form
2. Ticket Modification Form
3. Column Creation Form
4. Column Modificatoin Form
5. Board Creation Form
6. Board Modification Form

### Pages
- Board selection page
- Board Page
- Ticket Page if so desire but want to leave as popups

### Routes
/
- Homepage which is the board selection screen

/{boardTitle}
- Page home to the board with that particular title

/{ticketID}
- Page to check on the particular ticket
- Kinda overkill not really gonna use this much

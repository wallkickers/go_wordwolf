package domain

type Member struct {
    Id        int
    GroupId int
    AccountName string
    role Role
    voteTo int
    theme string
}

type Members []Member

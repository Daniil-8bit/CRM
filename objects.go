package main

type Opportunity struct {
	OppId      int
	OppNumber  int
	OppName    string
	oppOwner   Contact
	oppLead    Lead
	oppAccount Account
}

type Lead struct {
	LeadId      int
	LeadName    string
	leadOwner   Contact
	leadAccount Account
	LeadSource  string
}

type Contact struct {
	ContactId         int
	ContactSurname    string
	ContactName       string
	ContactMiddlename string
	contactCompany    *Account
	ContactPhone      string
	ContactEmail      string
	ContactJobTitle   string
	contactOwner      *Contact
}

type Account struct {
	accountId          int
	accountName        string
	accountFullname    string
	accountMainContact Contact
	accountOwner       Contact
	accountInn         uint64
	accountWeb         string
}

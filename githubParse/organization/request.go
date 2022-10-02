package organization

import (
	"fmt"
	"net/http"

	. "githubParse/githubParse/member"
	. "githubParse/githubParse/repository"
	"githubParse/githubParse/utils"
)

const OrganisationURL = utils.BaseURL + "/orgs/"
const OrganisationMembersURL = "/members"

func Information(client *http.Client, organizationName string) (*Organization, error) {
	organization, err := GetOrganisationsInfo(client, organizationName)
	if err != nil {
		return nil, err
	}
	members, err := GetOrganisationsMembers(client, organizationName)
	if err != nil {
		return nil, err
	}
	repositories, err := GetOrganisationsRepositories(client, organizationName)
	if err != nil {
		return nil, err
	}

	addOrganisationsMembers(organization, members)
	addOrganisationsRepositories(organization, repositories)

	return organization, nil

}

func GetOrganisationsInfo(client *http.Client, organizationName string) (*Organization, error) {
	res, err := utils.GetRequest(client, OrganisationURL+organizationName)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	organization := Organization{}
	err = utils.DecodeResponseJSON(res, &organization)
	if err != nil {
		return nil, fmt.Errorf("organization information decode issue: %s", err)
	}
	return &organization, nil

}

func GetOrganisationsMembers(client *http.Client, organizationName string) (*[]*Member, error) {
	res, err := utils.GetRequest(client, OrganisationURL+organizationName+OrganisationMembersURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var members []*Member
	//body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))
	err = utils.DecodeResponseJSON(res, &members)
	if err != nil {
		return nil, fmt.Errorf("organization members decode issue: %s", err)
	}

	return &members, nil
}

func addOrganisationsMembers(organisation *Organization, members *[]*Member) {
	var allMembers []string
	membersInformation := make(map[string]*Member)

	for _, member := range *members {
		membersInformation[member.Login] = member
		allMembers = append(allMembers, member.Login)
	}

	organisation.Members = membersInformation
	organisation.MemberNames = allMembers
}

func GetOrganisationsRepositories(client *http.Client, organizationName string) (*[]*Repository, error) {
	res, err := utils.GetRequest(client, OrganisationURL+organizationName+RepositoryURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var repositories []*Repository
	err = utils.DecodeResponseJSON(res, &repositories)
	if err != nil {
		return nil, fmt.Errorf("organization repositories decode issue: %s", err)
	}
	return &repositories, nil

}

func addOrganisationsRepositories(organisation *Organization, repositories *[]*Repository) {
	var allRepositories []string
	repositoriesInformation := make(map[string]*Repository)

	for _, repository := range *repositories {
		repositoriesInformation[repository.Name] = repository
		allRepositories = append(allRepositories, repository.Name)
	}

	organisation.Repositories = repositoriesInformation
	organisation.RepositoryNames = allRepositories
}

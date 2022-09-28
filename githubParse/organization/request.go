package organization

import (
	"fmt"
	"net/http"

	. "githubParse/githubParse/member"
	"githubParse/githubParse/utils"
)

const OrganisationURL = utils.BaseURL + "/orgs/"
const OrganisationMembersURL = "/members"

func Information(client *http.Client, organizationName string) (*Organization, error) {
	var allMembers []string
	membersInformation := make(map[string]*Member)

	organization, err := GetOrganisationsInfo(client, organizationName)
	if err != nil {
		return nil, err
	}
	members, err := GetOrganisationsMembers(client, organizationName)
	if err != nil {
		return nil, err
	}

	for _, member := range *members {
		membersInformation[member.Login] = member
		allMembers = append(allMembers, member.Login)
	}

	organization.Members = membersInformation
	organization.MemberNames = allMembers

	return organization, nil

}

func GetOrganisationsInfo(client *http.Client, organizationName string) (*Organization, error) {
	res, err := utils.GetRequest(client, OrganisationURL+organizationName)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	organization := Organization{}
	//body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))
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

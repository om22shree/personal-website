import streamlit as st
from PIL import Image


# --- PAGE CONFIGURATION ---
st.set_page_config(
    page_title="Om Shree",
    page_icon=":briefcase:",
    layout="wide",
    initial_sidebar_state="expanded",
)


# --- PATHS ---
PROFILE_PIC_PATH = "data/rick-morty.jpg"  
RESUME_PATH = "data/resume.pdf" 
DEGREE_PATH = "data/degree.pdf"
CKAD_PATH = "data/ckad.pdf"
SYSTEMDESIGN_PATH = "data/systemDesign.pdf"
GITOPS_PATH = "data/gitOpsEnterprise.pdf"


# --- HEADER SECTION ---
with st.container():
    col1, col2 = st.columns([1, 2])
    with col1:
        try:
            profile_pic = Image.open(PROFILE_PIC_PATH)
            st.image(profile_pic, width=500, caption="I like this a little too much")
        except FileNotFoundError:
            st.error(f"Profile picture not found at {PROFILE_PIC_PATH}. Please add it.")
            st.image("https://via.placeholder.com/230", width=500) # Placeholder

    with col2:
        st.title("Om says :wave:")
        st.subheader("Cloud Platform | DevOps | Data streaming")
        st.write(
            "Passionate about building scalable and efficient data streaming architectures using kubernetes and AWS"
        )
        if RESUME_PATH:
            try:
                with open(RESUME_PATH, "rb") as pdf_file:
                    PDFbyte = pdf_file.read()
                st.download_button(
                    label="📄 Download Resume",
                    data=PDFbyte,
                    file_name=RESUME_PATH.split('/')[-1], # Gets the file name
                    mime="application/octet-stream",
                )
            except FileNotFoundError:
                st.warning(f"Resume PDF not found at {RESUME_PATH}. Download button will not be shown.")


# --- ABOUT ME ---
with st.container():
    st.write("---")
    st.header("About Me")
    st.write(
        """
        Hello! I'm Om, a cloud engineer passionate about infrastructure, data streams and automation.
        I have a strong foundation in Kubernetes, Kafka, AWS, Terraform, Pulsar and enjoy solving complex problems
        and building innovative solutions.

        - 🔭 I’m currently working at FICO 
        - 🌱 I’m currently building data streaming platforms using apache kafka, pulsar, kubernetes, AWS and GH actions
        - 👯 I’m looking to collaborate on DevOps, data infrastructure, pipelines and streams
        - 💬 Ask me about FastAPI, Pulsar, Kafka, ArgoCD, AWS, DAPR, Kubernetes, KEDA and ELK stack
        - 📫 How to reach me: om22shree@gmail.com [trusty old-school email]
        - 😄 Pronouns: Him / His
        - ⚡ Fun fact: I'm a big fan of WarHammer 40k and daily-drive Fedora Linux
        """
    )


# --- EXPERIENCE ---
with st.container():
    st.write("---")
    st.header("Work Experience")
    st.subheader("Cloud Engineer 1 at FICO | July 2024 - Present")
    st.write(
        """
        - Automated deployment and patching of multiple production-grade **StreamNative Pulsar clusters** using nested **Helm charts**, **GitHub shared workflows**, **Kustomize**, **ArgoCD** and **Istio**
        - Implemented **autoscaling** for Pulsar broker and bookkeeper pods using **HPA**, **KEDA** and **prometheus**
        custom metrics
        - **On-call support**, operations, performance optimizations and monitoring improvements for production **Apache Kafka** clusters using **ELK stack**, **Grafana**, **AWS EC2**, **route53** and **JVM tuning**
        - Developed a **Kubernetes mutating webhook** to insert **DAPR** sidecars into applications based on envrionment specific configurations
        - **Promoted** to Cloud Engineer 1 from Cloud Engineer associate in June 2025
        """
    )
    st.subheader("Cloud Engineer Intern at FICO | Jun 2023 - June 2024")
    st.write(
        """
        - Developed enterprise wide **Multi AWS account** life-cycle management automation using **Terraform**, **Boto3**, **Python3** and **Jenkins**
        - Participated in **CSA STAR level 2** and **PCI attestation audit** - cloud infrastructure
        - Responsible for enhancements, automations, maintenance and **on-call** support for enterprise-wide **Route53** records, **IAM** policies and **AWS organization** components
        """
    )


# --- EDUCATION ---
with st.container():
    st.write("---")
    st.header("Education & Certifications")
    st.subheader("B. Tech - Information Technology | KIIT Bhubaneswar | 2020 - 2024")
    if RESUME_PATH:
        try:
            with open(DEGREE_PATH, "rb") as pdf_file:
                PDFbyte = pdf_file.read()
            st.download_button(
                label="📄 Download Degree",
                data=PDFbyte,
                file_name=  DEGREE_PATH.split('/')[-1],     
                mime="application/octet-stream",
            )
        except FileNotFoundError:
            st.warning(f"Degree certificate PDF not found at {DEGREE_PATH}. Download button will not be shown.")
    st.write(
        """
        - Relevant coursework: Machine Learning, Data Structures, Computer Networks, Cloud Computing, Operting Systems")
        - 9.2 CGPA
        - Particidpated in Global data science conclave
        - Participated in Zuno fellowship program
        """
    )
    st.subheader("CKAD - Certified Kubernetes Application Developer | 2024")
    if CKAD_PATH:
        try:
            with open(CKAD_PATH, "rb") as pdf_file:
                PDFbyte = pdf_file.read()
            st.download_button(
                label="📄 Download CKAD Certificate",
                data=PDFbyte,
                file_name=CKAD_PATH.split('/')[-1], 
                mime="application/octet-stream",
            )
        except FileNotFoundError:
            st.warning(f"CKAD certificate PDF not found at {CKAD_PATH}. Download button will not be shown.")
    
    st.subheader("GitOps Enterprise - level 3 | 2024")
    if GITOPS_PATH:
        try:
            with open(GITOPS_PATH, "rb") as pdf_file:
                PDFbyte = pdf_file.read()
            st.download_button(
                label="📄 Download GitOps Certificate",
                data=PDFbyte,
                file_name=GITOPS_PATH.split('/')[-1], 
                mime="application/octet-stream",
            )
        except FileNotFoundError:
            st.warning(f"GitOps certificate PDF not found at {GITOPS_PATH}. Download button will not be shown.")
    
    st.subheader("System Design - HLD & LLD | 2022")
    if SYSTEMDESIGN_PATH:
        try:
            with open(SYSTEMDESIGN_PATH, "rb") as pdf_file:
                PDFbyte = pdf_file.read()
            st.download_button(
                label="📄 Download System Design Certificate",
                data=PDFbyte,
                file_name=SYSTEMDESIGN_PATH.split('/')[-1], 
                mime="application/octet-stream",
            )
        except FileNotFoundError:
            st.warning(f"System Design certificate PDF not found at {SYSTEMDESIGN_PATH}. Download button will not be shown.")
    

# --- PROJECTS ---
with st.container():
    st.write("---")
    st.header("Projects")
    st.write("##") 

    # Project 1
    col1, col2 = st.columns([1, 2]) 
    with col1:
        try:
            project1_image = Image.open("data/airflowProject.png") 
            st.image(project1_image, caption="Airlfow Project sample image")
        except FileNotFoundError:
            st.image("https://via.placeholder.com/300x200?text=Project+1+Image", caption="Project 1 Placeholder")
    with col2:
        st.subheader("Batch Transaction Processor and analyzer")
        st.write(
            """
            **Description:** Provides a robust solution for processing and analyzing large volumes of batch transactions. It utilizes Apache Airflow for orchestration, PostgreSQL for storage, FastAPI for API endpoints & ELK stack for monitoring and logging.
            
            **Technologies Used:** Apache Airflow, PostgreSQL, FastAPI, Docker, Kubernetes, ELK Stack
            
            **Achievements:**
            - Successfully processed over 1 million transactions in a single batch
            - Reduced processing time through optimized DAGs and parallel task execution
            - Implemented comprehensive monitoring and alerting using ELK stack & grafana
            """
        )
        st.markdown("[View Source Code](https://github.com/om22shree/airflow-transaction-analyzer)") 

    st.write("---")

    # Project 2
    col1, col2 = st.columns([2, 1]) 
    with col1:
        st.subheader("Kubernetes MWC and controller for DAPR sidecar injection")
        st.write(
            """
            **Description:** This project involved creating a Kubernetes mutating webhook controller that automatically injects DAPR sidecars into applications based on environment-specific configurations. It also keeps a track of the injected sidecars and their configurations using a dynamic configmap.
            
            **Technologies Used:** Kubernetes, DAPR, Go, Helm

            **Achievements:**
            - Developed a mutating webhook controller in Go that dynamically injects DAPR sidecars into Kubernetes pods basis namespace configurations
            - Created a Helm chart to deploy the mutating webhook controller and manage its lifecycle
            - Enhanced application observability and resilience by integrating DAPR sidecars for service discovery, state management, and pub/sub messaging
            """
        )
        st.markdown("[View Source Code](https://github.com/yourusername/kubernetes-dapr-webhook)")
    with col2:
        try:
            project2_image = Image.open("data/kubernetesProject.png")
            st.image(project2_image, caption="Kubernetes project sample image")
        except FileNotFoundError:
            st.image("https://via.placeholder.com/200x300?text=Project+2+Image", caption="Kubernetes project sample image")


# --- SKILLS ---
with st.container():
    st.write("---")
    st.header("My Skills")
    col1, col2 = st.columns(2)
    with col1:
        st.subheader("Languages & Frameworks")
        st.write("""
        - Python
        - SQL
        - bash/Shell
        - Go
        - FastAPI
        - Linux
        """)
    with col2:
        st.subheader("Platform & Tooling")
        st.write("""
        - AWS (S3, EC2, Route53, LoadBalancer, EKS, Boto3)
        - Apache Kafka and Pulsar
        - Terraform
        - Kubernetes, Istio and DAPR
        - ArgoCD
        - GitHub Actions and Workflows
        - ELK stack
        """)


# --- CONTACT ---
with st.container():
    st.write("---")
    st.subheader("Get In Touch!")
    st.write("""
        - [LinkedIn](https://linkedin.com/in/om22shree)
        - [Gmail](mailto:om22shree@gmail.com)
        """)


# --- FOOTER ---
with st.container():
    st.markdown("© Om Shree 2025 | Made with [Streamlit](https://streamlit.io/)")